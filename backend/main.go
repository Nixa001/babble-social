package main

import (
	server "backend/app"
	"backend/utils"
)

func main() {
	// seed.InsertData(seed.DB)
	utils.ClearScreen()
	server := server.NewServer()
	server.Run()
}
