package handler

import (
	"backend/models"
	"backend/server/cors"
	"encoding/json"
	"io"
	"net/http"
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
			return
		}
		w.WriteHeader(http.StatusCreated)
		var user models.User
		err = json.Unmarshal(content, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(user)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
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
		if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
			return
		}
		json.NewEncoder(w).Encode(credentials)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
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
		json.NewEncoder(w).Encode(map[string]any{"message": "success"})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
