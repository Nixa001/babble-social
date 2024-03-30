package handler

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/service"
	"backend/server/ws"
	"backend/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign up handler")
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPost:
		body := r.Body
		content, err := io.ReadAll(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		}
		fmt.Println("content:", content)
		var formatedUser models.FormatedUser
		err = json.Unmarshal(content, &formatedUser)
		fmt.Println("FormatedUser:", formatedUser)
		if err != nil {
			fmt.Println("Invalid request", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		} else if err := utils.IsValidEmail(strings.TrimSpace(formatedUser.Email)); err != nil {
			fmt.Println("Invalid email", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email"})
			return
		} else if err := utils.VerifyName(strings.TrimSpace(formatedUser.First_name)); err != nil {
			fmt.Println("Invalid first name", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid first name"})
			return
		} else if err := utils.VerifyName(strings.TrimSpace(formatedUser.Last_name)); err != nil {
			fmt.Println("Invalid last name", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid last name"})
			return
		} else if err := utils.VerifyUsername(formatedUser.User_name); err != nil {
			fmt.Println("Invalid username", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid username"})
			return
		}
		formatedUser.User_type = "public"
		err = service.AuthServ.CreateUser(formatedUser)
		if err != nil {
			log.Println("Error creating user", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		user, err := service.AuthServ.UserRepo.GetUserByEmail(formatedUser.Email)
		if err != nil {
			log.Println("Handler Error getting user", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		session, err := service.AuthServ.CreateSession(user)
		if err != nil {
			log.Println("Error creating session", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		// var newEvent = ws.WSPaylaod{
		// 	From: user.Email,
		// 	Type: ws.WS_JOIN_EVENT,
		// 	Data: map[string]any{
		// 		"username":     user.Email,
		// 		"status":       "offline",
		// 		"unread_count": 0,
		// 	},
		// }
		// ws.WSHub.HandleEvent(newEvent)
		w.WriteHeader(http.StatusOK)
		log.Printf("User %s created successfully", user.Email)
		json.NewEncoder(w).Encode(map[string]any{"message": "success", "token": session.Token, "user": user})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
	// ...
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign in handler")
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
			fmt.Println("Invalid credentials 0", err)
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
		email, password := credentials["email"], credentials["password"]

		if utils.IsValidEmail(strings.TrimSpace(email)) != nil || utils.IsValidPassword(strings.TrimSpace(password)) != nil {
			fmt.Println("Invalid credentials 1", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 1"})
			return
		}
		user, err := service.AuthServ.CheckCredentials(email, password)
		if err != nil {
			fmt.Println("Invalid credentials 2", err)
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials 2"})
			return
		}

		session, err := service.AuthServ.CreateSession(user)
		if err != nil {
			fmt.Println("Error creating session", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		log.Println("User Connected", user.Email)
		json.NewEncoder(w).Encode(map[string]any{"message": "success", "token": session.Token, "user": user})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		token := r.Header.Get("Authorization")
		if strings.TrimSpace(token) == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		}
		user, err := service.AuthServ.UserRepo.GetUserByToken(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		err = service.AuthServ.RemoveSession(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		if c, ok := ws.WSHub.Clients.Load(user.Email); ok {
			client := c.(*ws.WSClient)
			ws.WSHub.UnRegisterChannel <- client
			var newEvent = ws.WSPaylaod{
				// From: client.Email,
				Type: ws.WS_DISCONNECT_EVENT,
				Data: nil,
			}
			ws.WSHub.HandleEvent(newEvent)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"message": "success"})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}

}

func VerifySessionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Verify session handler")
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	session, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"message": "success", "token": session.Token})

}
