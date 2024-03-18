package routes

import (
	"backend/server/handler"
	"net/http"
)

const (
	HOME_ENDPOINT   = "/"
	SIGNUP_ENDPOINT = "/auth/signup"
	SIGNIN_ENDPOINT = "/auth/signin"
	LOGOUT_ENDPOINT = "/auth/signout"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(HOME_ENDPOINT, handler.LoggingMiddleware(handler.IndexHandler))
	mux.HandleFunc(SIGNUP_ENDPOINT, handler.LoggingMiddleware(handler.SignUpHandler))
	mux.HandleFunc(SIGNIN_ENDPOINT, handler.LoggingMiddleware(handler.SignInHandler))
	mux.HandleFunc(LOGOUT_ENDPOINT, handler.LoggingMiddleware(handler.AuthorizeMiddleware(handler.SignOutHandler)))

	return mux
}
