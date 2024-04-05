package handler

import (
	"backend/server/cors"
	"encoding/json"
	"fmt"
	"net/http"
)

type Test struct {
	Content string `json:"message"`
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	switch r.Method {
	case "POST":
		var test Test
		err := json.NewDecoder(r.Body).Decode(&test)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		fmt.Println(test.Content)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(test)
	default:

	}
}
