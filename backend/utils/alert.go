package utils

import (
	"backend/models"
	"encoding/json"
	"net/http"
)

func Alert(w http.ResponseWriter, msg models.Errormessage) {
	response := map[string]interface{}{}
	response["msg"] = msg.Msg
	response["status"] = msg.StatusCode
	response["type"] = msg.Type
	response["display"] = msg.Display
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(msg.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func AlertPostData(w http.ResponseWriter, msg models.WResponse) {
	response := map[string]interface{}{}
	response["msg"] = msg.Msg
	response["status"] = msg.StatusCode
	response["type"] = msg.Type
	response["display"] = msg.Display
	response["data"] = msg.Data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(msg.StatusCode)
	json.NewEncoder(w).Encode(response)
}
