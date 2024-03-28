package handler

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/handler/session"
	"backend/utils/seed"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch r.Method {
	case "POST":
		var credentials models.Credentials
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		stm := `SELECT * FROM users WHERE user_name = ? OR email = ?`
		query, err := seed.DB.Query(stm, credentials.Email, credentials.Email)
		if err != nil {
			fmt.Println("Error getting credentials")
			return
		}
		defer query.Close()
		var log = false
		var user models.User
		for query.Next() {
			err := query.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Username, &user.Gender, &user.Email, &user.Password, &user.UserType, &user.BirthDate, &user.Avatar, &user.AboutMe)
			if err != nil {
				fmt.Println("Error getting credentials", err)
				return
			}
			if (user.Username == credentials.Email || user.Email == credentials.Email) && (credentials.Password == user.Password) {
				log = true
				user = models.User{
					ID:        user.ID,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
					Username:  user.Username,
					Gender:    user.Gender,
					Email:     user.Email,
					UserType:  user.UserType,
					BirthDate: user.BirthDate,
					Avatar:    user.Avatar,
					AboutMe:   user.AboutMe,
				}
			}
			var reponse = map[string]interface{}{}
			if log {
				reponse = map[string]interface{}{
					"status":  log,
					"connect": true,
					"message": "connection established",
					"page":    "home",
					"user":    user,
					"token":   "123",
				}
				fmt.Println("userid: ", user.ID)
				session.CreateSession(w, r, user.ID)

			}

			responsJson, err := json.Marshal(reponse)
			if err != nil {
				fmt.Println(err.Error())
			}
			w.Write(responsJson)
		}
	default:
	}

}
