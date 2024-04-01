package handler

import (
	"backend/server/cors"
	"backend/server/ws"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, world!")
	cors.SetCors(&w)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur de mise a niveau de la connexion", err)
		return
	}
	ws.Init()
	ws.WSHub.AddClient(conn, "mail@a.com")

	// go func() {
	// 	for {
	// 		_, message, err := conn.ReadMessage()
	// 		if err != nil {
	// 			log.Println("Erreur lors de la lecture du message:", err)
	// 			return
	// 		}
	// 		log.Println("Message reçu:", string(message))
	// 	}
	// }()

	fmt.Println("Ws is running")
	// _ = conn.Close()
}
