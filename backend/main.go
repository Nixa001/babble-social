package main

import (
	"backend/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}
