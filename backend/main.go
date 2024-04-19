package main

import (
	server "backend/server"
)

func main() {
	// seed.InsertData(seed.DB)
	// utils.ClearScreen()
	server := server.NewServer()
	server.Run()
}
