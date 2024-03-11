package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	WS_JOIN_EVENT       = "join-event"
	WS_DISCONNECT_EVENT = "disconnect-event"
	WS_ADD_FEED_POST    = "add-feed-post"
	WS_ADD_GROUP_POST   = "add-group-post"
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
	To   []string
}

type Hub struct {
	Clients           *sync.Map
	RegisterChannel   chan *WSClient
	UnRegisterChannel chan *WSClient
	SSE               chan WSPaylaod
}

var WSHub *Hub

func Init() {
	WSHub = NewHub()
	go WSHub.listen()
}

func NewHub() *Hub {
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
			for _, to := range eventPayload.To {
				if client.Firstname == to {
					client.OutgoingMsg <- eventPayload
				}
			}
			return true
		})
	/*-------------------------------------------------------------
	--------------------------------------------------------------*/
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			for _, to := range eventPayload.To {
				if client.Firstname == to {
					client.OutgoingMsg <- eventPayload
				}
			}
			return true
		})
	case WS_ADD_FEED_POST:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if eventPayload.To[0] == "all" {
				client.OutgoingMsg <- eventPayload
			} else {
				for _, to := range eventPayload.To {
					if client.Firstname == to {
						client.OutgoingMsg <- eventPayload
					}
				}
			}
			return true
		})
	case WS_ADD_GROUP_POST:
		//! handle group posts here
		fmt.Println("in group...")
	}
}

func (wsHub *Hub) AddClient(coon *websocket.Conn, firstname string) {
	client := &WSClient{
		Firstname:   firstname,
		WSCoon:      coon,
		OutgoingMsg: make(chan interface{}),
	}
fmt.Println("client is here :", client)
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
		fmt.Println("reading")
		_, message, err := client.WSCoon.ReadMessage()
		if err != nil {
			WSHub.UnRegisterChannel <- client

			var newEvent = WSPaylaod{
				From: client.Firstname,
				Type: WS_DISCONNECT_EVENT,
				Data: nil,
			}
			WSHub.HandleEvent(newEvent)
			panic(err)
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
