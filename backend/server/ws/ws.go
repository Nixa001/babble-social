package ws

import (
	"backend/utils/seed"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	WS_JOIN_EVENT             = "join-event"
	WS_DISCONNECT_EVENT       = "disconnect-event"
	WS_IDRECEIVER_EVENT       = "id-receiver-event"
	WS_MESSAGEUSER_EVENT      = "message-user-event"
	WS_MESSAGEGROUP_EVENT     = "message-group-event"
	WS_IDGROUP_RECEIVER_EVENT = "idGroup-receiver-event"
)

type WSClient struct {
	Email       string
	WSCoon      *websocket.Conn
	OutgoingMsg chan interface{}
}
type WSPaylaod struct {
	From string
	Type string
	Data interface{}
	To   []string
}
type Hub struct {
	Clients           *sync.Map
	RegisterChannel   chan *WSClient
	UnRegisterChannel chan *WSClient
	SSE               chan WSPaylaod
}

var WSHub *Hub

func init() {
	WSHub = newHub()
	go WSHub.listen()
}
func newHub() *Hub {
	return &Hub{
		Clients:           &sync.Map{},
		RegisterChannel:   make(chan *WSClient),
		UnRegisterChannel: make(chan *WSClient),
		SSE:               make(chan WSPaylaod),
	}
}
func (h *Hub) listen() {
	for {
		select {
		case client := <-h.RegisterChannel:
			h.Clients.Store(client.Email, client)
			log.Printf("Client %s connected\n", client.Email)
		case client := <-h.UnRegisterChannel:
			if _, ok := h.Clients.Load(client.Email); ok {
				h.Clients.Delete(client.Email)
				close(client.OutgoingMsg)
				log.Printf("Client %s disconnected\n", client.Email)
			}
		case message := <-h.SSE:
			h.HandleEvent(message)
		}
	}
}
func (h *Hub) HandleEvent(eventPayload WSPaylaod) {
	switch eventPayload.Type {
	case WS_JOIN_EVENT:
		fmt.Println("tous les client", h.Clients)
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			for _, name := range eventPayload.To {
				if client.Email == name {
					client.OutgoingMsg <- eventPayload
					break // Exit the loop once a match is found
				}
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			fmt.Println("2", client)
			for _, name := range eventPayload.To {
				if client.Email == name {
					client.OutgoingMsg <- eventPayload
					break // Exit the loop once a match is found
				}
			}
			return true
		})
	case WS_IDRECEIVER_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// for _, name := range eventPayload.To {
			// 	if client.Email == name {
			client.OutgoingMsg <- eventPayload
			// break // Exit the loop once a match is found
			// 	}
			// }
			return true
		})
	case WS_IDGROUP_RECEIVER_EVENT:
		h.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			// if client.Email == eventPayload.From {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	case WS_MESSAGEUSER_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Email == eventPayload.From {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	case WS_MESSAGEGROUP_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// for _, name := range eventPayload.To {
			// 	if client.Email == name {
			client.OutgoingMsg <- eventPayload
			// break // Exit the loop once a match is found
			// 	}
			// }
			return true
		})
	}
}
func (wsHub *Hub) AddClient(coon *websocket.Conn, Email string) {
	client := &WSClient{
		Email:       Email,
		WSCoon:      coon,
		OutgoingMsg: make(chan interface{}),
	}
	go client.messageReader()
	go client.messageWriter()
	wsHub.RegisterChannel <- client
	// var newEvent = WSPaylaod{
	// 	From: client.Email,
	// 	Type: WS_JOIN_EVENT,
	// 	Data: nil,
	// }
	// wsHub.HandleEvent(newEvent)
	// return
}
func (client *WSClient) messageReader() {
	for {
		_, message, err := client.WSCoon.ReadMessage()
		if err != nil {
			WSHub.UnRegisterChannel <- client
			var newEvent = WSPaylaod{
				From: client.Email,
				Type: WS_DISCONNECT_EVENT,
				Data: nil,
			}
			WSHub.HandleEvent(newEvent)
			return
		}
		var payload map[string]interface{}
		err = json.Unmarshal(message, &payload)
		if err != nil {
			return
		}
		eventType := payload["type"].(string)
		// EvenType le type d'evenement socket
		switch eventType {

		case WS_JOIN_EVENT:
			var newEvent = WSPaylaod{
				From: client.Email,
				Type: WS_JOIN_EVENT,
				Data: "New client joined",
			}
			WSHub.HandleEvent(newEvent)

		case WS_IDRECEIVER_EVENT:
			data, _ := seed.SelectMsgBetweenUsers(seed.DB, 1, 2)
			wsEvent := WSPaylaod{
				From: "Ndiba",
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(wsEvent)
			return

		case WS_IDGROUP_RECEIVER_EVENT:
			data, _ := seed.GetGroupMessage(seed.DB, 1)
			wsEvent := WSPaylaod{
				From: "",
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(wsEvent)
			return

		case WS_MESSAGEUSER_EVENT:

			err := seed.InsertMessage(seed.DB, 1, 2, payload["data"].(string), "2000-01-01")
			if err != nil {
				fmt.Println(err)
			}
			data := "communication between User"
			msEvent := WSPaylaod{
				From: "",
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(msEvent)
			return

		case WS_MESSAGEGROUP_EVENT:
			data := "communication between Group"
			GPEvent := WSPaylaod{
				From: "",
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(GPEvent)
			return
		}

	}
}
func (client *WSClient) messageWriter() {
	for {
		select {
		case message := <-client.OutgoingMsg:
			data, err := json.Marshal(message)
			if err != nil {
				return
			}
			err = client.WSCoon.WriteMessage(websocket.TextMessage, data)
			fmt.Println(string(data))
			if err != nil {
				return
			}
		}
	}
}
