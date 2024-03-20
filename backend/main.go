package main

import (
	"backend/server"
	utils "backend/utils"
)

func main() {
	utils.ClearScreen()
	server := server.NewServer()

	
	server.Run()
}
