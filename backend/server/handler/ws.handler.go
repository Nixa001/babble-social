package handler

import (
	"log"
	"net/http"

	"backend/server/ws"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("socket request from %s", r.RemoteAddr)
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
		r.Body.Close()

	}()
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrading error: ", err)
		return
	}
	ws.Init()
	ws.WSHub.AddClient(con, "mass") // todo: add user id here
}
