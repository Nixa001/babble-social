package handler

import (
	"backend/server/cors"
	"net/http"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// ...
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// ...
}

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// ...
}
