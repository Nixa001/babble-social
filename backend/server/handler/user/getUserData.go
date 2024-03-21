package user

import (
	"backend/models"
	"backend/server/cors"
	"backend/utils/seed"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

var userId int = 1

type RespUser struct {
	User      models.User   `json:"user"`
	Followers []models.User `json:"followers"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var respUser RespUser
	cors.SetCors(&w)
	var db = seed.CreateDB()
	defer db.Close()

	respUser.User, _ = GetUserData(db, userId)

	respUser.Followers, _ = GetFollowers(db, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respUser)
}

func GetUserData(db *sql.DB, userID int) (models.User, error) {
	var user models.User

	query := "SELECT * FROM users WHERE id =?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return user, fmt.Errorf("err execute user query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.UserType, &user.BirthDate, &user.Avatar, &user.AboutMe)
		if err != nil {
			return user, fmt.Errorf("err scan group data: %w", err)
		}
	}

	if err := rows.Err(); err != nil {
		return user, fmt.Errorf("err read all groups results: %w", err)
	}

	// fmt.Println(user)
	return user, nil
}

func GetFollowers(db *sql.DB, userID int) ([]models.User, error) {
	var followers []models.User

	query := "SELECT user_id_follower FROM users_followers WHERE user_id_followed = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("err execute all groups query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, fmt.Errorf("err scan group data: %w", err)
		}
		user, err := GetUserData(db, id)
		if err != nil {
			return nil, fmt.Errorf("err execute all groups query: %w", err)
		}
		followers = append(followers, user)
	}

	return followers, nil
}
