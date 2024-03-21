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
			h.Clients.Store(client.Email, client)
			log.Printf("Client %s connected\n", client.Email)
		case client := <-h.UnRegisterChannel:
			if _, ok := h.Clients.Load(client.Email); ok {
				h.Clients.Delete(client.Email)
				close(client.OutgoingMsg)
				log.Printf("Client %s disconnected\n", client.Firstname)
			}
		case message := <-h.SSE:
			h.HandleEvent(message)

		}
	}
}

func (h *Hub) HandleEvent(eventPayload WSPaylaod) {
	fmt.Println("in handle event...")
	switch eventPayload.Type {
	case "GET":
		fmt.Println("👨‍💻 in test case...")
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			//for _, t := range eventPayload.To {
			//if client.Firstname == to {
			client.OutgoingMsg <- eventPayload
			//}
			//}
			return true
		})
	case WS_JOIN_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Firstname == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	/*-------------------------------------------------------------
	--------------------------------------------------------------*/
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Firstname == eventPayload.To {
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

func (wsHub *Hub) AddClient(coon *websocket.Conn, Email string) {
	client := &WSClient{
		Email:       Email,
		WSCoon:      coon,
		OutgoingMsg: make(chan interface{}),
	}
	fmt.Println("client is here :", client)
	go client.messageReader()
	go client.messageWriter()

	wsHub.RegisterChannel <- client

	var newEvent = WSPaylaod{
		From: client.Email,
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
				From: client.Email,
				Type: WS_DISCONNECT_EVENT,
				Data: nil,
			}
			WSHub.HandleEvent(newEvent)
			fmt.Println("exit")
			return
		}
		var payload map[string]interface{}
		err = json.Unmarshal(message, &payload)
		if err != nil {
			return
		}
		eventType := payload["type"].(string)
		fmt.Println("type is here => ", eventType)
		fmt.Println("payload is here => ", payload)
		//	i:=0
		//	for {
		wsEvent := WSPaylaod{
			From: client.Email,
			Type: eventType,
			Data: payload,
		}
		WSHub.HandleEvent(wsEvent)
		// i++
		// fmt.Printf("test %v done\n", i)
		// time.Sleep(3 * time.Second) // TODO : make it dynamic
		//	}
		//fmt.Println("done reading and sent to hdle event!")
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
