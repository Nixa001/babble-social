package routes

import (
	"backend/server/handler"
	"backend/server/handler/groups"
	"backend/server/handler/user"
	"net/http"
)

const (
	HOME_ENDPOINT          = "/"
	SIGNUP_ENDPOINT        = "/auth/signup"
	SIGNIN_ENDPOINT        = "/auth/signin"
	LOGOUT_ENDPOINT        = "/auth/signout"
	POST_ENDPOINT          = "/post"
	COMMENT_ENDPOINT       = "/comment"
	CREATE_GROUP_ENDPOINT  = "/group/creategroup"
	CREATE_EVENT_ENDPOINT  = "/group/createEvent"
	POST_GROUP_ENDPOINT    = "/group/postgroup"
	USER_ENDPOINT          = "/userInfo"
	JOINED_GROUPS_ENDPOINT = "/groups_joined"
	GETGROUPS_ENDPOINT     = "/groups"
	GETGROUP_ENDPOINT      = "/groups/group"
	WS_ENDPOINT            = "/ws"
	SERVE_ASSETS           = "/uploads/"
)

func Route() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(HOME_ENDPOINT, handler.IndexHandler)
	mux.HandleFunc(SIGNUP_ENDPOINT, handler.SignUpHandler)
	mux.HandleFunc(SIGNIN_ENDPOINT, handler.SignInHandler)
	mux.HandleFunc(LOGOUT_ENDPOINT, handler.AuthorizeMiddleware(handler.SignOutHandler))
	mux.HandleFunc(POST_ENDPOINT, handler.POSTHandler)
	mux.HandleFunc(COMMENT_ENDPOINT, handler.COMMENTHandler)
	mux.HandleFunc(CREATE_GROUP_ENDPOINT, groups.CreateGroupHandler)
	mux.HandleFunc(CREATE_EVENT_ENDPOINT, groups.CreateEventHandler)
	mux.HandleFunc(POST_GROUP_ENDPOINT, groups.PostGroupHandler)
	mux.HandleFunc(GETGROUPS_ENDPOINT, groups.GetGroups)
	mux.HandleFunc(WS_ENDPOINT, handler.WSHandler)
	mux.HandleFunc(GETGROUP_ENDPOINT, groups.GetGroup)
	mux.HandleFunc(USER_ENDPOINT, user.GetUser)
	mux.HandleFunc(SERVE_ASSETS, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	return mux
}
