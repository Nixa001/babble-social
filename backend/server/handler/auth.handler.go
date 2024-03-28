package handler

import (
	"backend/models"
	"backend/server/cors"
	"backend/server/service"
	"backend/server/ws"
	"backend/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
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
		var user models.User
		err = json.Unmarshal(content, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		} else if err := utils.IsValidEmail(strings.TrimSpace(user.Email)); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email"})
			return
		} else if err := utils.VerifyName(strings.TrimSpace(user.First_name)); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid first name"})
			return
		} else if err := utils.VerifyName(strings.TrimSpace(user.Last_name)); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid last name"})
			return
		} else if err := utils.VerifyUsername(user.User_name); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid username"})
			return
		}
		user.User_type = "public"
		err = service.AuthServ.CreateUser(&user)
		if err != nil {
			log.Println("Error creating user", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		token := utils.GenerateToken()
		err = service.AuthServ.SessRepo.CreateSession(&models.Session{Token: token, User_id: user.Id, Expiration: utils.GenerateExpirationTime()})
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
		json.NewEncoder(w).Encode(map[string]any{"message": "success", "token": token, "user": user})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
	}
	// ...
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case http.MethodPost:
		credentials := make(map[string]string, 2)
		content, err := io.ReadAll(r.Body)
		if err == nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}

		err = json.Unmarshal(content, &credentials)
		if err == nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}
		email, password := credentials["email"], credentials["password"]

		if utils.IsValidEmail(strings.TrimSpace(email)) != nil || utils.IsValidPassword(strings.TrimSpace(password)) != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email"})
			return
		}
		user, err := service.AuthServ.CheckCredentials(email, password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}
		err = service.AuthServ.RemExistingSession(strconv.Itoa(user.Id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		token := utils.GenerateToken()
		err = service.AuthServ.SessRepo.CreateSession(&models.Session{User_id: user.Id, Token: token, Expiration: utils.GenerateExpirationTime()})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"message": "success", "token": token, "user": user})
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
		//recupere le user à partir du token
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
				From: client.Mail,
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
	user, err := service.AuthServ.UserRepo.GetUserByToken(session.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{"message": "success", "user": user})

}
