package models

type Session struct {
	Token      string `json:"token"`
	User_id    int    `json:"user_id"`
	Expiration string `json:"expiration"`
}
