package routes

import (
	"backend/server/handler"
	"net/http"
)

const (
	SIGNUP_ENDPOINT     = "/auth/signup"
	SIGNIN_ENDPOINT     = "/auth/signin"
	LOGOUT_ENDPOINT     = "/auth/signout"
	VERIF_SESS_ENDPOINT = "/auth/session"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(SIGNUP_ENDPOINT, (handler.SignUpHandler))
	mux.HandleFunc(SIGNIN_ENDPOINT, (handler.SignInHandler))
	mux.HandleFunc(LOGOUT_ENDPOINT, (handler.AuthorizeMiddleware(handler.SignOutHandler)))
	mux.HandleFunc(VERIF_SESS_ENDPOINT, (handler.VerifySessionHandler))

	return mux
}
