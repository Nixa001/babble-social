package groups

import (
	"backend/models"
	"backend/server/cors"
	joingroup "backend/server/handler/groups/joinGroup"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const userID int = 1

func GetGroups(w http.ResponseWriter, r *http.Request) {

	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()
	joinedGroups, err := getJoinedGroups(db, userID)
	if err != nil {
		return
	}

	allGroups, err := getAllGroups(db)
	if err != nil {
		return
	}

	filteredGroups, groups := filterGroups(db,joinedGroups, allGroups)
	var Groups [][]models.Group
	Groups = append(Groups, groups)
	Groups = append(Groups, filteredGroups)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Groups)
}

func getJoinedGroups(db *sql.DB, userID int) ([]int, error) {
	var joinedGroupIDs []int

	query := "SELECT group_id FROM group_followers WHERE user_id = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute joined group query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("failed to scan joined group data: %w", err)
		}
		joinedGroupIDs = append(joinedGroupIDs, id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read joined groups results: %w", err)
	}

	return joinedGroupIDs, nil
}

func getAllGroups(db *sql.DB) ([]models.Group, error) {
	var allGroups []models.Group

	query := "SELECT * FROM groups"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var group models.Group
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			return nil, fmt.Errorf("failed to scan group data: %w", err)
		}
		allGroups = append(allGroups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read all groups results: %w", err)
	}

	return allGroups, nil
}

func filterGroups(db *sql.DB, joined []int, all []models.Group) ([]models.Group, []models.Group) {
	var filteredGroups []models.Group
	var Groups []models.Group
	for _, group := range all {
		isJoined := false
		for _, id := range joined {
			if id == group.ID {
				isJoined = true
				break
			}
		}
		if !isJoined && group.ID_User_Create != userID {
			check, state := joingroup.CheckJoinNotification(group.ID_User_Create, userID, group.ID,db)
			if check !=0 &&state==0{
				group.State ="disable"
			}

			filteredGroups = append(filteredGroups, group)
		} else {
			groupId := strconv.Itoa(group.ID)

			group.Href = "/home/groups/group_id=" + groupId
			Groups = append(Groups, group)
		}
		// fmt.Println(group)
	}
	return filteredGroups, Groups
}
// func checkNotif(db *sql.DB, groupID, userID int) (bool, error) {
// 	var exists bool
// 	err := db.QueryRow("SELECT EXISTS(SELECT id FROM notifications WHERE notification_type = ?)", groupName).Scan(&exists)
// 	if err != nil {
// 		return false, err
// 	}
// 	return exists, nil
// }