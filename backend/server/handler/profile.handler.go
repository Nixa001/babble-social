package handler

import (
	"backend/server/cors"
	"backend/server/service"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	cors.SetCors(&w)

	switch r.Method {
	case http.MethodOptions:
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodGet:
		userIDStr := r.URL.Query().Get("id")
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
			return
		}
		log.Println("User ID:", userID)
		// get user by id
		user, err := service.AuthServ.UserRepo.GetUserById(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Failed to get user:", err)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get user"})
			return
		}
		// // get user's posts
		// posts, err := service.PostServ.FetchPostGroupByUserID(userID)
		// if err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user's posts"})
		// 	return
		// }
		// get user's followers
		followers, err := service.AuthServ.GetFollowersByID(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Failed to get user's followers:", err)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get user's followers"})
			return
		}
		// get user's followings
		followings, err := service.AuthServ.GetFollowingsByID(userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to get user's followings"})
			return
		}
		// return user info, posts, followers, followings
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"user": user, "followers": followers, "followings": followings})
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func FollowHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Sign up handler")
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch r.Method {
	case http.MethodPost:
		idinquery, err := url.QueryUnescape(r.URL.Query().Get("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		}
		idinquery = strings.ReplaceAll(idinquery, " ", "+")
		idinquery = strings.TrimSpace(idinquery)
		idtofollow, err := strconv.Atoi(idinquery)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		}
		session, err := service.AuthServ.VerifyToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
			return
		}
		id := session.User_id
		err = service.AuthServ.FollowUser(id, idtofollow)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Followed"})
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
