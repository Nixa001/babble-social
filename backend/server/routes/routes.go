package routes

import (
	"backend/server/handler"
	groups "backend/server/handler/groups"
	"net/http"
)

const (
	HOME_ENDPOINT          = "/"
	SIGNUP_ENDPOINT        = "/auth/signup"
	SIGNIN_ENDPOINT        = "/auth/signin"
	LOGOUT_ENDPOINT        = "/auth/signout"
	CREATE_GROUP_ENDPOINT  = "/group/creategroup"
	JOINED_GROUPS_ENDPOINT = "/groups_joined"
	GETGROUPS_ENDPOINT     = "/groups"
	GETGROUP_ENDPOINT     = "/groups/group"
	WS_ENDPOINT            = "/socket"
	SERVE_ASSETS           = "/uploads/"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(HOME_ENDPOINT, handler.IndexHandler)
	mux.HandleFunc(SIGNUP_ENDPOINT, handler.SignUpHandler)
	mux.HandleFunc(SIGNIN_ENDPOINT, handler.SignInHandler)
	mux.HandleFunc(LOGOUT_ENDPOINT, handler.AuthorizeHandler(handler.SignOutHandler))
	mux.HandleFunc(CREATE_GROUP_ENDPOINT, groups.CreateGroupHandler)
	mux.HandleFunc(WS_ENDPOINT, handler.WSHandler)
	mux.HandleFunc(GETGROUPS_ENDPOINT, groups.GetGroups)
	mux.HandleFunc(GETGROUP_ENDPOINT, groups.GetGroup)
	mux.HandleFunc(SERVE_ASSETS, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	return mux
}
