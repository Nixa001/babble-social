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
	Email       string
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
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
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
	var newEvent = WSPaylaod{
		From: client.Email,
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
		wsEvent := WSPaylaod{
			From: client.Email,
			Type: eventType,
			Data: payload,
		}
		WSHub.HandleEvent(wsEvent)
	}
}
func (client *WSClient) messageWriter() {
	for message := range client.OutgoingMsg {
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