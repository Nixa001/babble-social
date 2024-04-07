package handler

import (
	"backend/server/cors"
	"backend/server/service"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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
		credentials := make(map[string]string, 2)
		content, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Invalid credentials 0 :", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}

		err = json.Unmarshal(content, &credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 0"})
			return
		}
		followerIDStr, followingIDStr := credentials["follower_id"], credentials["following_id"]
		followerID, err := strconv.Atoi(followerIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid follower ID"})
			return
		}
		followingID, err := strconv.Atoi(followingIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid following ID"})
			return
		}
		err = service.AuthServ.FollowUser(followerID, followingID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to follow"})
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

func UnfollowHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch r.Method {
	case http.MethodPost:
		credentials := make(map[string]string, 2)
		content, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Invalid credentials 0 :", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}

		err = json.Unmarshal(content, &credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 0"})
			return
		}
		followerIDStr, followingIDStr := credentials["follower_id"], credentials["following_id"]
		followerID, err := strconv.Atoi(followerIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid follower ID"})
			return
		}
		followingID, err := strconv.Atoi(followingIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid following ID"})
			return
		}
		err = service.AuthServ.UnfollowUser(followerID, followingID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to unfollow"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Unfollowed"})
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func SwitchProfileType(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	switch r.Method {
	case http.MethodPost:
		credentials := make(map[string]string, 2)
		content, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Invalid credentials 0 :", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}

		err = json.Unmarshal(content, &credentials)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 0"})
			return
		}
		userIDStr, profileType := credentials["user_id"], credentials["profile_type"]
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid user ID"})
			return
		}
		err = service.AuthServ.UpdateProfileType(userID, profileType)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to update profile type"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Profile type updated"})
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
