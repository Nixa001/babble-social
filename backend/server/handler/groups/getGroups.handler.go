package groups

import (
	"backend/server/cors"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Group struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ID_User_Create int    `json:"id_user_create"`
	Avatar         string `json:"image"`
	Creation_Date  string `json:"creation_date"`
	Href           string `json:"href"`
}

const userID int = 1


func GetGroups(w http.ResponseWriter, r *http.Request) {
	
	cors.SetCors(&w)
	var db = seed.CreateDB()
	joinedGroups, err := getJoinedGroups(db, userID)
	if err != nil {
		return
	}
	
	allGroups, err := getAllGroups(db)
	if err != nil {
		return
	}
	
	filteredGroups, groups := filterGroups(joinedGroups, allGroups)
	var Groups [][]Group
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

func getAllGroups(db *sql.DB) ([]Group, error) {
	var allGroups []Group

	query := "SELECT * FROM groups"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var group Group
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

func filterGroups(joined []int, all []Group) ([]Group, []Group) {
	var filteredGroups []Group
	var Groups []Group
	for _, group := range all {
		isJoined := false
		for _, id := range joined {
			if id == group.ID {
				isJoined = true
				break
			}
		}
		if !isJoined && group.ID_User_Create != userID {
			filteredGroups = append(filteredGroups, group)
			} else {
			group.Href = "/home/groups/group/"
			Groups = append(Groups, group)
		}
		// fmt.Println(group)
	}
	return filteredGroups, Groups
}
