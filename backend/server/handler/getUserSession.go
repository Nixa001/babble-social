package handler

import (
	"backend/models"
	"backend/server/cors"
	"backend/utils/seed"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUserSession(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	//Extraire le tolen de la requete
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token is required", http.StatusBadRequest)
		return
	}

	//Recuperation de l"id du user

	var userID int

	err := seed.DB.QueryRow("SELECT user_id FROM sessions WHERE token = ?", token).Scan(&userID)
	if err != nil {
		http.Error(w, "User ID  not found", http.StatusNotFound)
		return
	}

	var userSession models.User
	err = seed.DB.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(&userSession.Id, &userSession.First_name, &userSession.Last_name, &userSession.User_name, &userSession.Gender, &userSession.User_type, &userSession.About_me, &userSession.Birth_date, &userSession.Email, &userSession.Avatar, &userSession.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userSession)
}
