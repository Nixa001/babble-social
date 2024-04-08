package handler

import (
	"backend/server/cors"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// ...
}
func AuthorizeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.SetCors(&w)
		next.ServeHTTP(w, r)
	}
}

func RequestValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.SetCors(&w)
		next.ServeHTTP(w, r)
	}
}
