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
	POST_ENDPOINT = "/post"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(HOME_ENDPOINT, handler.IndexHandler)
	mux.HandleFunc(SIGNUP_ENDPOINT, handler.SignUpHandler)
	mux.HandleFunc(SIGNIN_ENDPOINT, handler.SignInHandler)
	mux.HandleFunc(LOGOUT_ENDPOINT, handler.AuthorizeHandler(handler.SignOutHandler))
	mux.HandleFunc(POST_ENDPOINT, handler.POSTHandler)

	return mux
}
