package ws

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	WS_JOIN_EVENT       = "join-event"
	WS_DISCONNECT_EVENT = "disconnect-event"
)

type WSClient struct {
	Firstname   string
	WSCoon      *websocket.Conn
	OutgoingMsg chan interface{}
}

type WSPaylaod struct {
	From string
	Type string
	Data interface{}
	To   string
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
			h.Clients.Store(client.Firstname, client)
			log.Printf("Client %s connected\n", client.Firstname)
		case client := <-h.UnRegisterChannel:
			if _, ok := h.Clients.Load(client.Firstname); ok {
				h.Clients.Delete(client.Firstname)
				close(client.OutgoingMsg)
				log.Printf("Client %s disconnected\n", client.Firstname)
			}
		case message := <-h.SSE:
			h.HandleEvent(message)

		}
	}
}

func (h *Hub) HandleEvent(eventPayload WSPaylaod) {
	switch eventPayload.Type {
	case WS_JOIN_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Firstname == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Firstname == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	}
}

func (wsHub *Hub) AddClient(coon *websocket.Conn, firstname string) {
	client := &WSClient{
		Firstname:   firstname,
		WSCoon:      coon,
		OutgoingMsg: make(chan interface{}),
	}

	go client.messageReader()
	go client.messageWriter()

	wsHub.RegisterChannel <- client

	var newEvent = WSPaylaod{
		From: client.Firstname,
		Type: WS_JOIN_EVENT,
		Data: nil,
	}

	wsHub.HandleEvent(newEvent)

}

func (client *WSClient) messageReader() {
	for {
		_, message, err := client.WSCoon.ReadMessage()
		if err != nil {
			WSHub.UnRegisterChannel <- client

			var newEvent = WSPaylaod{
				From: client.Firstname,
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

		wsEvent := WSPaylaod{
			From: client.Firstname,
			Type: eventType,
			Data: payload,
		}

		WSHub.HandleEvent(wsEvent)
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
			if err != nil {
				return
			}

		}
	}
}
