package main

import (
	server "backend/app"
)

func main() {
	// seed.InsertData(seed.DB)
	// utils.ClearScreen()
	server := server.NewServer()
	server.Run()
}
