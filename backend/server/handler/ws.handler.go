package handler

import (
	"backend/server/service"
	"backend/server/ws"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	session, err := service.AuthServ.VerifyToken(r)
	if err != nil {
		//TODO handle error here
		return
	}
	coon, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//TODO handle error here
		fmt.Println("Error to ws")
	}
	user, err := service.AuthServ.UserRepo.GetUserById(session.User_id)
	if err != nil {
		fmt.Println("internal Server")
	}
	ws.WSHub.AddClient(coon, user.Email)
}
