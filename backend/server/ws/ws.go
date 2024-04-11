package ws

import (
	"backend/database"
	"backend/models"
	"backend/server/handler/groups/events"
	joingroup "backend/server/handler/groups/joinGroup"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

/*
* Le bloc `const` définit deux constantes `WS_JOIN_EVENT` et `WS_DISCONNECT_EVENT` avec des valeurs de
chaîne `"join-event"` et `"disconnect-event"` respectivement. Ces constantes sont utilisées pour
représenter différents types d'événements pouvant survenir dans le système de communication
WebSocket. Le `WS_JOIN_EVENT` est généralement utilisé lorsqu'un client rejoint la connexion
WebSocket, tandis que `WS_DISCONNECT_EVENT` est utilisé lorsqu'un client se déconnecte de la
connexion WebSocket. Ces constantes aident à maintenir une représentation cohérente et claire des
événements dans toute la base de code.
*/
const (
	WS_JOIN_EVENT       = "join-event"
	WS_DISCONNECT_EVENT = "disconnect-event"
)

/*
Le type WSClient représente un client WebSocket avec un mail, une connexion WebSocket et un canal
de message sortant.
@property {string} Mail - La propriété `Mail` dans la structure `WSClient` est un champ de
type chaîne qui stocke le mail du client associé à la connexion WebSocket.
@property WSCoon - La structure `WSClient` représente un client WebSocket. Voici un aperçu de ses
propriétés :
@property OutgoingMsg - La propriété `OutgoingMsg` dans la structure `WSClient` est un canal qui
peut être utilisé pour envoyer des messages ou des données à envoyer via la connexion WebSocket
représentée par le champ `WSCoon`. Ce canal permet la communication asynchrone des données qui
doivent être envoyées par le client.
*/
type WSClient struct {
	Mail        string
	WSCoon      *websocket.Conn
	OutgoingMsg chan interface{}
}

/*
La structure WSPayload définit une structure de données pour les charges utiles WebSocket contenant
des informations sur l'expéditeur, le type, les données et le destinataire.
@property {string} From - La propriété « From » dans la structure « WSPayload » représente
l'expéditeur de la charge utile. Il contient généralement des informations sur l'origine ou la
source des données transmises.
@property {string} Type - La propriété `Type` dans la structure `WSPayload` représente le type de
données envoyées dans la charge utile WebSocket. Il peut s'agir d'un type de message, d'un type
d'événement ou de toute autre catégorisation permettant de traiter les données du côté récepteur.
@property Data - La propriété `Data` dans la structure `WSPayload` est de type `interface{}`. Cela
signifie qu'il peut contenir des valeurs de n'importe quel type de données. Il s'agit d'un type
générique qui peut être utilisé pour stocker différents types de données, tels que des chaînes, des
entiers, des structures ou même des données personnalisées.
@property {string} To - La propriété « To » dans la structure « WSPayload » représente le
destinataire ou la destination des données utiles. Il spécifie le destinataire prévu du message ou
des données envoyées.
*/
type WSPaylaod struct {
	From string
	Type string
	Data interface{}
	To   string
}

/*
Le type `Hub` dans Go représente un hub pour la gestion des clients WebSocket et des canaux de
communication.
@property Clients - La propriété `Clients` dans la structure `Hub` est un pointeur vers un
`sync.Map`. Cette carte est généralement utilisée pour stocker et gérer les clients WebSocket
connectés au hub. Le type `sync.Map` est une implémentation de carte sécurisée en termes de
concurrence fournie par la bibliothèque standard Go. Il permet une simultanéité sûre
@property RegisterChannel - La propriété `RegisterChannel` dans la structure `Hub` est un canal
utilisé pour enregistrer de nouveaux clients WebSocket. Lorsqu'un nouveau client se connecte au
serveur WebSocket, une référence au client est envoyée via ce canal pour être enregistrée dans le
`Hub`.
@property UnRegisterChannel - La propriété `UnRegisterChannel` dans la structure `Hub` est un canal
utilisé pour envoyer des pointeurs `WSClient` pour indiquer qu'un client doit être désenregistré du
hub. Ce canal est probablement utilisé pour communiquer avec le hub lorsqu'un client doit être
supprimé ou déconnecté.
@property SSE - La propriété `SSE` dans la structure `Hub` est un canal utilisé pour envoyer les
données `WSPayload`. Il permet au « Hub » de communiquer avec les clients connectés en envoyant des
charges utiles d'événements envoyés par le serveur (SSE).
*/
type Hub struct {
	Clients           *sync.Map
	RegisterChannel   chan *WSClient
	UnRegisterChannel chan *WSClient
	SSE               chan WSPaylaod
}

var WSHub *Hub

/*
	La fonction `init` initialise un nouveau hub WebSocket et commence à écouter les messages entrants

dans une goroutine distincte.
*/
func Init() {
	WSHub = newHub()
	go WSHub.listen()
}

/*
	La fonction newHub renvoie une nouvelle instance de Hub avec des canaux initialisés et un sync.Map

pour gérer les clients WebSocket.
*/
func newHub() *Hub {
	return &Hub{
		Clients:           &sync.Map{},
		RegisterChannel:   make(chan *WSClient),
		UnRegisterChannel: make(chan *WSClient),
		SSE:               make(chan WSPaylaod),
	}
}

/*
La fonction `func (h *Hub) Listen()` est la boucle d'événements principale pour le hub WebSocket. Il
écoute en permanence les événements entrants provenant de différents canaux à l'aide d'une
instruction « select ». Voici une répartition de ce qu'il fait pour chaque cas :
*/
func (h *Hub) listen() {
	for {
		select {
		case client := <-h.RegisterChannel:
			h.Clients.Store(client.Mail, client)
			log.Printf("Client %s connected\n", client.Mail)
		case client := <-h.UnRegisterChannel:
			if _, ok := h.Clients.Load(client.Mail); ok {
				h.Clients.Delete(client.Mail)
				close(client.OutgoingMsg)
				log.Printf("Client %s disconnected\n", client.Mail)
			}
		case message := <-h.SSE:
			h.HandleEvent(message)

		}
	}
}

/* La méthode `HandleEvent` dans la structure `Hub` est responsable du traitement de différents types
d'événements WebSocket en fonction du champ `Type` de la structure `WSPaylaod` passé en argument.
*/

func (h *Hub) HandleEvent(eventPayload WSPaylaod) {
	switch eventPayload.Type {
	case WS_JOIN_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Mail == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case WS_DISCONNECT_EVENT:
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Mail == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "Welcome":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Mail == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "JoinGroup":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			// if client.Mail == eventPayload.To {
			client.OutgoingMsg <- eventPayload
			// }
			return true
		})
	case "NotGoingEvent":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Mail == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "GoingEvent":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Mail == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "Desplayed Events":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Mail == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "SuggestFriend":
		h.Clients.Range(func(key, value interface{}) bool {
			client := value.(*WSClient)
			if client.Mail == eventPayload.To {
				client.OutgoingMsg <- eventPayload
			}
			return true
		})
	case "notification":
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

/* La fonction `func (wsHub *Hub) AddClient(coon *websocket.Conn, Mail string)` est responsable de
l'ajout d'un nouveau client au hub WebSocket. Voici un aperçu de ce qu'il fait :
*/

func (wsHub *Hub) AddClient(coon *websocket.Conn, Mail string) {
	client := &WSClient{
		Mail:        Mail,
		WSCoon:      coon,
		OutgoingMsg: make(chan interface{}),
	}

	go client.messageReader()
	go client.messageWriter()

	wsHub.RegisterChannel <- client

	var newEvent = WSPaylaod{
		From: client.Mail,
		Type: WS_JOIN_EVENT,
		Data: nil,
	}

	wsHub.HandleEvent(newEvent)

}

/*
La fonction `func (client *WSClient) messageReader()` est une méthode définie sur la structure

`WSClient` dans Go. Cette méthode est chargée de lire les messages de la connexion WebSocket
associée au client.
*/
func (client *WSClient) messageReader() {
	Db := database.NewDatabase()
	defer Db.Close()
	for {
		_, message, err := client.WSCoon.ReadMessage()
		if err != nil {
			WSHub.UnRegisterChannel <- client

			var newEvent = WSPaylaod{
				From: client.Mail,
				Type: WS_DISCONNECT_EVENT,
				Data: nil,
			}
			WSHub.HandleEvent(newEvent)
			return
		}
		var payload map[string]interface{}
		fmt.Println("Message received", string(message))
		err = json.Unmarshal(message, &payload)
		if err != nil {
			return
		}
		eventType := payload["type"].(string)
		fmt.Println("message", eventType)
		wsEvent := WSPaylaod{
			From: client.Mail,
			Type: eventType,
			Data: payload,
		}

		fmt.Println("WsEvent", wsEvent.Data)

		switch wsEvent.Type {
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
			idUserConnect := 1
			err = joingroup.InsertNotification(int(groupeId), typeNotification, idUserConnect, Db)
			if err != nil {
				fmt.Println("Error inserting", err.Error())
			}
			fmt.Println("Notification joined added to database")

			dataSend := struct {
				IdGroup int    `json:"id_group"`
				Button  string `json:"button"`
			}{
				IdGroup: int(groupeId),
				Button:  "Desable",
			}

			wsEvent = WSPaylaod{
				From: client.Mail,
				Type: eventType,
				Data: dataSend,
				To:   "Adimine group",
			}
			WSHub.HandleEvent(wsEvent)

			// notification := joingroup.ListNotification(idUserConnect, Db)
			// fmt.Println("notification ", notification)
			// if notification != nil {
			// 	dataSend := struct {
			// 		IdGroup int                   `json:"id_group"`
			// 		Message []models.Notification `json:"message"`
			// 		Type    string                `json:"type"`
			// 		To      int                   `json:"to"`
			// 		Button  string                `json:"button"`
			// 	}{
			// 		IdGroup: int(groupeId),
			// 		Message: notification,
			// 		Type:    "Notification",
			// 		To:      idUserConnect,
			// 		Button:  "Desable",
			// 	}

			// 	wsEvent = WSPaylaod{
			// 		From: "",
			// 		Type: eventType,
			// 		Data: dataSend,
			// 		To:   "Adimine group",
			// 	}
			// 	fmt.Println("wsEvent: ", wsEvent)
			// 	WSHub.HandleEvent(wsEvent)
			// }

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
			// fmt.Println("eventId: ", event_id)
			// dataSend := struct {
			// 	Group_id int    `json:"id_group"`
			// 	Event_id int    `json:"button"`
			// 	Message  string `json:"message"`
			// 	Type     string ` json:"type"`
			// }{
			// 	Group_id: int(groupeId),
			// 	Event_id: int(event_id),
			// 	Message:  "Delete event for user",
			// 	Type:     "NotGoingEvent",
			// }

			// wsEvent = WSPaylaod{
			// 	From: client.Mail,
			// 	Type: eventType,
			// 	Data: dataSend,
			// 	To:   "Adimine group",
			// }
			// WSHub.HandleEvent(wsEvent)

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
			event_id, ok := parseData["event_id"].(float64)
			if !ok {
				fmt.Println("Erreur de recuperation de donnee")
				return
			}
			err = events.JoinEvent(1, int(groupeId), int(event_id), Db)
			if err != nil {
				log.Fatal(err.Error())
			}

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
			idUserConnect := 1
			group_id, err := strconv.Atoi(groupeID)
			if err != nil {
				log.Fatal(err.Error())

			}
			fmt.Println("WS -- ", group_id, typeNotification, idUserConnect)
			err = joingroup.InsertNotification(group_id, typeNotification, idUserConnect, Db)
			if err != nil {
				fmt.Println("Error inserting notification ", err.Error())
				return
			}
		case "notification":
			idUserConnect := 1
			notification := joingroup.ListNotification(idUserConnect, Db)
			if notification != nil {
				dataSend := struct {
					Message []models.Notification `json:"message"`
					Type    string                `json:"type"`
					To      int                   `json:"to"`
				}{
					Message: notification,
					Type:    "Notification",
					To:      idUserConnect,
				}

				wsEvent = WSPaylaod{
					From: "",
					Type: eventType,
					Data: dataSend,
					To:   "Admin group",
				}
				fmt.Println("Notification ", wsEvent)
				WSHub.HandleEvent(wsEvent)
			}
		case "ResponceNotification":
			fmt.Println("--- Notification responce ---")

		}

	}
}

/*
La fonction `func (client *WSClient) messageWriter()` est une méthode définie sur la structure
`WSClient` dans Go. Cette méthode est responsable de l'écriture des messages sur la connexion
WebSocket associée au client. Voici un aperçu de ce qu'il fait :
*/
func (client *WSClient) messageWriter() {
	// Créer un canal pour les messages entrants
	// incomingMsg := make(chan string)

	// // Routine pour lire les messages entrants et les afficher
	// go func() {
	// 	for {
	// 		message := <-incomingMsg
	// 		fmt.Println("Incoming message:", message) // Affiche le message reçu sur le terminal
	// 	}
	// }()

	// Boucle principale pour l'envoi de messages sortants
	for {
		select {
		case message := <-client.OutgoingMsg:
			// Convertir le message en JSON
			data, err := json.Marshal(message)
			if err != nil {
				fmt.Println("Error marshaling message:", err)
				return
			}

			// Envoyer le message sur la connexion WebSocket
			err = client.WSCoon.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				// Gérer l'erreur `ErrCloseSent`
				if errors.Is(err, websocket.ErrCloseSent) {
					fmt.Println("Connexion WebSocket fermée par le serveur")
					return
				} else {
					fmt.Println("Error writing message to WebSocket:", err)
					return
				}
			}

			// Envoyer également le message sur le canal des messages entrants
			// incomingMsg <- string(data)
		}
	}
}

// Cette fonction envoie un message à un client WebSocket avec les données de réponse fournies.
func SendMessage(client *WSClient, reponsData interface{}) {
	payload := WSPaylaod{
		From: client.Mail,
		Type: "JoinGroup",
		Data: reponsData,
		To:   client.Mail,
	}
	fmt.Println("=========Payload: ", payload.Data)
	client.OutgoingMsg <- payload
}

// query = "SELECT id, event_id, user_id, group_id, description, event_date FROM event_joined INNER JOIN event ON event_joined.event_id = event.id WHERE event_joined.user_id = ?"
