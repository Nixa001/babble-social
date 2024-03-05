package main

import (
	"backend/utils"
	"backend/utils/seed"
	"fmt"
	"log"
	"net/http"
)

func main() {
	utils.ClearScreen()
	seed.CreateTable(seed.DB)
	seed.InsertData(seed.DB)
	fmt.Println("http://localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
