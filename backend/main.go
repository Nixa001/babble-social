package main

import (
	"backend/server"
	"backend/utils"
)

func main() {
	// seed.InsertData(seed.DB)
	utils.ClearScreen()
	server := server.NewServer()
	server.Run()
}