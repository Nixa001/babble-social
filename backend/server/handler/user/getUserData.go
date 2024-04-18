package user

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/handler/groups"
	"backend/server/service"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// var userId int = 1

type RespUser struct {
	User      models.User    `json:"user"`
	Followers []models.User  `json:"followers"`
	Groups    []models.Group `json:"groups"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var respUser RespUser
	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()
	session, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		fmt.Println("Invalid token")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}

	userId := session.User_id

	respUser.User, _ = GetUserData(db, userId)
	idJoinedGroups, err1 := groups.GetJoinedGroups(db, userId)
	respUser.Groups, err = GetGroupData(db, idJoinedGroups)
	if err != nil || err1 != nil {
		fmt.Println("errr when getGroups")
		return
	}
	respUser.Followers, _ = GetFollowers(db, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respUser)
}

func GetGroupData(db *sql.DB, groupsId []int) ([]models.Group, error) {
	var groups []models.Group
	for _, idGroup := range groupsId {
		var group models.Group
		query := "SELECT * FROM groups WHERE id = ?"
		err := db.QueryRow(query, idGroup).Scan(&group.ID, &group.Name, &group.Description, &group.ID_User_Create, &group.Avatar, &group.Creation_Date)
		if err != nil {
			fmt.Println("errr ici")
			return []models.Group{}, fmt.Errorf("err scan user data: %w", err)
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func GetUserData(db *sql.DB, userID int) (models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE id = ?"
	err := db.QueryRow(query, userID).Scan(&user.Id, &user.First_name, &user.Last_name, &user.User_name, &user.Gender, &user.Email, &user.Password, &user.User_type, &user.Birth_date, &user.Avatar, &user.About_me)
	if err != nil {
		fmt.Println("errr ici")
		return models.User{}, fmt.Errorf("err scan user data: %w", err)
	}
	return user, nil
}

func GetFollowers(db *sql.DB, userID int) ([]models.User, error) {
	var followers []models.User

	query := "SELECT user_id_follower FROM users_followers WHERE user_id_followed = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		fmt.Println("err@ ici")
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	var followerIDs []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}
		followerIDs = append(followerIDs, id)
	}

	query = "SELECT user_id_followed FROM users_followers WHERE user_id_follower = ?"
	rows, err = db.Query(query, userID)
	if err != nil {
		fmt.Println("err@ ici")
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}

		if !contains(followerIDs, id) {

			followerIDs = append(followerIDs, id)
		}

	}

	for _, followerID := range followerIDs {
		user, err := GetUserData(db, followerID)
		if err != nil {
			return nil, fmt.Errorf("err execute all groups query: %w", err)
		}
		followers = append(followers, user)
	}

	return followers, nil
}

func contains(arr []int, value int) bool {
	for _, element := range arr {
		if element == value {
			return true
		}
	}
	return false
}
