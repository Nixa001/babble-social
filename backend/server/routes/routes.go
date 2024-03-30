package routes

import (
	"backend/server/handler"
	"net/http"
)

const (
	HOME_ENDPOINT    = "/"
	SIGNUP_ENDPOINT  = "/auth/signup"
	SIGNIN_ENDPOINT  = "/auth/signin"
	LOGOUT_ENDPOINT  = "/auth/signout"
	MESSAGE_ENDPOINT = "/message"
	WS_ENDPOINT      = "/ws"
	LOGIN_ENDPOINT   = "/login"
	VERIF_SESS_ENDPOINT = "/auth/session"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(HOME_ENDPOINT, handler.IndexHandler)
	mux.HandleFunc(MESSAGE_ENDPOINT, handler.MessageHandler)
	mux.HandleFunc(WS_ENDPOINT, handler.WSHandler)
	mux.HandleFunc(LOGOUT_ENDPOINT, handler.AuthorizeHandler(handler.SignOutHandler))
	mux.HandleFunc(LOGIN_ENDPOINT, handler.LoginHandler)
	mux.HandleFunc(SIGNUP_ENDPOINT, (handler.SignUpHandler))
	mux.HandleFunc(SIGNIN_ENDPOINT, (handler.SignInHandler))
	mux.HandleFunc(LOGOUT_ENDPOINT, (handler.AuthorizeMiddleware(handler.SignOutHandler)))
	mux.HandleFunc(VERIF_SESS_ENDPOINT, (handler.VerifySessionHandler))

	return mux
}
