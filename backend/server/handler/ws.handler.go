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

	conn, err := upgrader.Upgrade(w, r, nil) // Utilisez l'Upgrader configuré
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer conn.Close()
	ws.WSHub.AddClient(conn, "Dicks")
	// return
	// for {
	// 	_, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// fmt.Println(string(message))

	// 	var msg Message
	// 	err = json.Unmarshal(message, &msg)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// fmt.Println(msg)
	// 	data1, err := seed.SelectMsgBetweenUsers(seed.DB, 1, 2)
	// 	if err != nil {
	// 		fmt.Println("Error")
	// 	}
	// 	data := "HTTP/1.0 200 OK LOGIN"
	// 	// data1 := "Content-Type: application/json MESSAGE"
	// 	if msg.Type == "join" {
	// 		conn.WriteJSON(data)
	// 	}
	// 	if msg.Type == "success" {
	// 		err := conn.WriteJSON(data1)
	// 		if err != nil {
	// 			fmt.Println("data error")
	// 		}
	// 	}
	// 	if msg.Type == "idUser-receiver-event" {
	// 		conn.WriteJSON("idUser-receiver")
	// 	}
	// 	if msg.Type == "idGroup-receiver-event" {
	// 		conn.WriteJSON("idGroup-receiver")
	// 	}

	// }
}
