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
	VERIF_SESS_ENDPOINT = "/auth/session"
	GET_USER_SESS_ENDPOINT = "/auth/usersessions"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(MESSAGE_ENDPOINT, handler.MessageHandler)
	mux.HandleFunc(WS_ENDPOINT, handler.WSHandler)
	mux.HandleFunc(SIGNUP_ENDPOINT, (handler.SignUpHandler))
	mux.HandleFunc(SIGNIN_ENDPOINT, (handler.SignInHandler))
	mux.HandleFunc(LOGOUT_ENDPOINT, (handler.AuthorizeMiddleware(handler.SignOutHandler)))
	mux.HandleFunc(VERIF_SESS_ENDPOINT, (handler.VerifySessionHandler))
	mux.HandleFunc(GET_USER_SESS_ENDPOINT, (handler.GetUserSession))

	return mux
}
