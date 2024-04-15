package notification

import (
	"backend/server/cors"
	joingroup "backend/server/handler/groups/joinGroup"
	"encoding/json"
	"net/http"
)

func GetNotification(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// db := database.NewDatabase()
	listeNotificaton := joingroup.ListNotification(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listeNotificaton)
}
