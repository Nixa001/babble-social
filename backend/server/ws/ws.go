package ws

import (
	// "backend/server/handler"

	"backend/database"
	"backend/server/handler/groups/events"
	joingroup "backend/server/handler/groups/joinGroup"
	"backend/server/service"
	"backend/server/service"
	"backend/utils/seed"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	WS_JOIN_EVENT             = "join-event"
	WS_DISCONNECT_EVENT       = "disconnect-event"
	WS_IDRECEIVER_EVENT       = "id-receiver-event"
	WS_MESSAGEUSER_EVENT      = "message-user-event"
	WS_MESSAGEGROUP_EVENT     = "message-group-event"
	WS_IDGROUP_RECEIVER_EVENT = "idGroup-receiver-event"
	WS_NAVBAR_MESSAGE         = "message-navbar"
)

type WSClient struct {
	Email        string
	WSCoon       *websocket.Conn
	OutgoingMsg  chan interface{}
	SessionToken string
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
			// if _, ok := h.Clients.Load(client.Email); ok {
			h.Clients.Delete(client.Email)
			close(client.OutgoingMsg)
			log.Printf("Client %s disconnected\n", client.Email)
			// }
		case message := <-h.SSE:
			h.HandleEvent(message)
		}
	}
}

var groupid int

func (h *Hub) HandleEvent(eventPayload WSPaylaod) {
	switch eventPayload.Type {
	case WS_JOIN_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email != eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_NAVBAR_MESSAGE:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)

			if client.Email == eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}

			return true
		})
	case WS_IDRECEIVER_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_IDGROUP_RECEIVER_EVENT:
		h.Clients.Range(func(key, value any) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.From {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_MESSAGEUSER_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.From || client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_MESSAGEGROUP_EVENT:
		clientFollowers, err := seed.GetFollowerGroup(seed.DB, groupid)
		if err != nil {
			fmt.Println("error", err)
		}
		fmt.Println("clientFollowers", clientFollowers)
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// Vérification si le client suit le groupe
			for _, follower := range clientFollowers {
				if client.Email == follower.Email {
					client.OutgoingMsg <- eventPayload
					break
				}
			}
			return true
		})

		//===============================================

	case "Welcome":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Email == eventPayload.To {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	case "JoinGroup":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "NotGoingEvent":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "GoingEvent":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "Desplayed Events":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "SuggestFriend":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "follow":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Email == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})	case "notification":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Mail == eventPayload.To {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	case "ResponceNotification":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Mail != eventPayload.To {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	}

}

func (wsHub *Hub) AddClient(coon *websocket.Conn, Email string, sessionToken string, r *http.Request) {
	client := &WSClient{
		Email:        Email,
		WSCoon:       coon,
		OutgoingMsg:  make(chan interface{}),
		SessionToken: sessionToken,
	}
	go client.messageReader(r)
	go client.messageWriter()
	wsHub.RegisterChannel <- client
	var newEvent = WSPaylaod{
		From: client.Email,
		Type: WS_JOIN_EVENT,
		Data: "New client joined",
	}
	wsHub.HandleEvent(newEvent)
}
func (client *WSClient) messageReader(r *http.Request) {

	userIdConnected, err := service.AuthServ.VerifyToken(r)
	fmt.Println("User connected : ==== ", userIdConnected)
	if err != nil {
		fmt.Println("verify token error", err)
	}

	date := time.Now().Format("2006-01-02T15:04:05")
	Db := database.NewDatabase()
	defer Db.Close()
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

		// EvenType le type d'evenement socket
		switch eventType {

		case WS_NAVBAR_MESSAGE:
			listGroup, _ := seed.GetGroup(seed.DB, int(payload["data"].(float64)))
			listUser, err := seed.ListeUsers(seed.DB, int(payload["data"].(float64)))
			if err != nil {
				fmt.Println("error:", err)
			} else if len(listUser) == 0 {
				fmt.Println("La liste des utilisateurs est vide")
			} else {

				data := []interface{}{listUser[0], listGroup}

				var newEvent = WSPaylaod{
					From: client.Email,
					Type: eventType,
					Data: data, // Assigner le tableau à Data
				}
				WSHub.HandleEvent(newEvent)
			}

		case WS_IDRECEIVER_EVENT:
			clickedUserId, ok := payload["data"].(map[string]interface{})["clickedUserId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'clickedUserId'")
				return
			}
			sessionUserId, ok := payload["data"].(map[string]interface{})["sessionUserId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}

			data, _ := seed.SelectMsgBetweenUsers(seed.DB, int(sessionUserId), int(clickedUserId))
			wsEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(wsEvent)

		case WS_IDGROUP_RECEIVER_EVENT:

			groupId, ok := payload["data"].(map[string]interface{})["idgroup"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'clickedUserId'")
				return
			}

			data, _ := seed.GetGroupMessage(seed.DB, int(groupId))
			wsEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: data,
			}
			WSHub.HandleEvent(wsEvent)

		case WS_MESSAGEUSER_EVENT:

			message, ok := payload["data"].(map[string]interface{})["message"].(string)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'message'")
				return
			}
			sendId, ok := payload["data"].(map[string]interface{})["sendId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}
			receiverID, ok := payload["data"].(map[string]interface{})["receiverId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}

			err := seed.InsertMessage(seed.DB, int(sendId), int(receiverID), message, date)
			if err != nil {
				fmt.Println(err)
			}
			mes, errr := seed.GetLastMessage(seed.DB, message)
			if errr != nil {
				fmt.Println(errr)
			}
			clientTo, _ := seed.GetUserById(seed.DB, int(receiverID))
			msEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: mes,
				To:   clientTo.Email,
			}
			WSHub.HandleEvent(msEvent)

		case WS_MESSAGEGROUP_EVENT:

			message, ok := payload["data"].(map[string]interface{})["message"].(string)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'message'")
				return
			}
			sendId, ok := payload["data"].(map[string]interface{})["sendId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}
			groupReceiverID, ok := payload["data"].(map[string]interface{})["receiverId"].(float64)
			if !ok {
				fmt.Println("Erreur lors de l'accès à la clé 'sendId'")
				return
			}
			groupid = int(groupReceiverID)
			err := seed.InsertGroupMessage(seed.DB, int(sendId), int(groupReceiverID), message, date)
			if err != nil {
				fmt.Println("error", err)
			}
			lastmsg, errr := seed.GetLastGroupMessage(seed.DB, message)
			if errr != nil {
				fmt.Println("error", err)

			}
			GPEvent := WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: lastmsg,
			}
			WSHub.HandleEvent(GPEvent)

			//==================================

		case "JoinGroup":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			groupeId, ok := parseData["groupId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}

			typeNotification, ok := parseData["type"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			err = joingroup.InsertNotification(int(groupeId), typeNotification, userIdConnected.User_id, Db)
			if err != nil {
				fmt.Println("Error inserting", err.Error())
			}
			fmt.Println("Notification joined added to database")

			idAdminGroup, err := joingroup.RecupeIdAdminGroup(int(groupeId), Db)
			if err != nil {
				fmt.Println("Erreur lors de la recuperation de l'id de l'administrateur du groupe ", err)
			}
			dataSend := struct {
				IdGroup int    `json:"id_group"`
				Button  string `json:"button"`
			}{
				IdGroup: int(groupeId),
				Button:  "Desable",
			}

			wsEvent = WSPaylaod{
				From: client.Email,
				Type: eventType,
				Data: dataSend,
				To:   string(idAdminGroup),
			}
			WSHub.HandleEvent(wsEvent)

		case "NotGoingEvent":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			groupeId, ok := parseData["groupId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			event_id, ok := parseData["event_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}

			err = events.NotJoinEvent(int(groupeId), 1, int(event_id), Db)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println("eventId: ", event_id)

		case "GoingEvent":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			groupeId, ok := parseData["groupId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}

			token, ok := parseData["token"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation du token join event")
				return
			}
			userID, err := service.AuthServ.VerifyTokenStr(token)
			if err != nil {
				fmt.Println("Erreur de recuperation de donnee VerifyTokenStr")
				return
			}

			event_id, ok := parseData["event_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			err = events.JoinEvent(userIdConnected.User_id, int(groupeId), int(event_id), Db)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(groupeId)
			fmt.Println(event_id)

			// fmt.Println("aaa", Db.Ping())
			// err = going.InsertNotification(int(groupeId), Db)
		case "SuggestFriend":
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			// fmt.Println("Parse json = ", parseData["id_group"])
			_, ok := parseData["userId"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}

			groupeID, ok := parseData["id_group"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			// fmt.Println("idgroup", groupeID)
			typeNotification, ok := parseData["type"].(string)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			// idUserConnect := 1
			group_id, err := strconv.Atoi(groupeID)
			if err != nil {
				log.Fatal(err.Error())

			}
			// fmt.Println("WS -- ", group_id, typeNotification, userIdConnected.User_id)
			err = joingroup.InsertNotification(group_id, typeNotification, userIdConnected.User_id, Db)
			if err != nil {
				fmt.Println("Error inserting notification ", err.Error())
				return
			}
		// case "notification":
		// 	fmt.Println("Notif")
		// 	notification := joingroup.ListNotification(userIdConnected.User_id, Db)
		// 	fmt.Println("mmmmmmmmmmmmmmm = ", notification)
		// 	if notification != nil {
		// 		dataSend := struct {
		// 			Message []models.Notification `json:"message"`
		// 			Type    string                `json:"type"`
		// 			To      int                   `json:"to"`
		// 		}{
		// 			Message: notification,
		// 			Type:    "Notification",
		// 			To:      userIdConnected.User_id,
		// 		}

		// 		wsEvent = WSPaylaod{
		// 			From: "",
		// 			Type: eventType,
		// 			Data: dataSend,
		// 			To:   "Admin group",
		// 		}
		// 		fmt.Println("Notification ", wsEvent)
		// 		WSHub.HandleEvent(wsEvent)
		// 	}
		case "ResponceNotification":
			fmt.Println("--- Notification responce ---")
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}

			// fmt.Println("Parse json = ", parseData["id_group"])
			id_user_sender, ok := parseData["id_user_sender"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			id_user_receiver, ok := parseData["id_user_receiver"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}

			groupIDSTr, ok := parseData["idGroup"].(string)
			groupeID, err := strconv.Atoi(groupIDSTr)
			fmt.Println("idgroup", groupeID)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			if err != nil {
				fmt.Println("cannot convert id group to int")
				return
			}
		case "follow":
			fmt.Println("follow")
			fmt.Println("wsEvent", wsEvent)
			jsonData, err := json.Marshal(wsEvent.Data)
			if err != nil {
				fmt.Println("Erreur de conversion en json", err)
				return
			}
			var parseData map[string]interface{}
			if err := json.Unmarshal(jsonData, &parseData); err != nil {
				fmt.Println("Erreur de conversion en json", err)
			}
			followerId, ok := parseData["follower_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			fmt.Println("followerId", followerId)
			followedId, ok := parseData["followed_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			fmt.Println("followedId", followedId)
			fmt.Println("database", Db.Ping())
			err = joingroup.InsertFollowNotification(int(followerId), int(followedId), Db)
			if err != nil {
				fmt.Println("Error inserting", err.Error())
			}
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
			// fmt.Println(string(data))
			if err != nil {
				return
			}
		}
	}
}
