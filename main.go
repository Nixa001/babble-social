package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Start the server
	fmt.Println("Server listening on: \nhttp://localhost:8000/")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Erreur au niveau du demarage du server... ", err.Error())
	}
}
