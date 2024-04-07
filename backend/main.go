package main

import (
	"backend/server"
	"backend/utils"
	"backend/utils/seed"
	"fmt"
)

func main() {
	// seed.InsertData(seed.DB)
	utils.ClearScreen()
	groups, err := seed.GetGroup(seed.DB, 1)
	if err != nil {
		fmt.Println("error getting group", err)
	}
	fmt.Println("groups:", groups)
	server := server.NewServer()
	server.Run()
}
