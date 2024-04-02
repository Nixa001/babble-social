package handler

import (
	"backend/server/cors"
	"backend/server/ws"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Autoriser toutes les origines pour le développement
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w) // Configurez les en-têtes CORS pour les réponses HTTP
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	userSession := UserInSession(w, r)

	conn, err := upgrader.Upgrade(w, r, nil) // Utilisez l'Upgrader configuré
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()
	ws.WSHub.AddClient(conn, userSession.Email)

}
