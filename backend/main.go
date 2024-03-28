package main

import (
	"backend/server"
	"backend/utils"
	"backend/utils/seed"
	"fmt"
)

func main() {
	utils.ClearScreen()
	seed.CreateTable(seed.DB)
	// message, err := seed.SelectMsgBetweenUsers(seed.DB, 1, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("message entre deux user: ", message)
	// for _, msg := range message {
	// 	fmt.Println(msg.ID, "===", msg.MessageContent, "===", msg.Date)
	// }
	// users, err := seed.SelectFollowersAndFollowing(seed.DB, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("--------------------------------")
	// fmt.Println("liste followers")
	// for _, user := range users {
	// 	fmt.Println(user.Firstname)
	// }
	// fmt.Println("--------------------------------")
	// fmt.Println("liste users in order")
	// msg, err := seed.ListeUsers(seed.DB, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("all messages", msg[0])
	// for _, user := range msg[0] {
	// 	fmt.Println(user.Firstname)
	// }
	// seed.InsertData(seed.DB)
	// err := seed.InsertMessage(seed.DB, 1, 2, "4 message", "2000-01-01")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// nsg, err := seed.GetGroupMessage(seed.DB, 1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(nsg)
	fmt.Println("http://localhost:8080")
	server := server.NewServer()
	server.Run()
}
