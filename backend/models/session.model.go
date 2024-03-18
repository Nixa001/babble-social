package models

type Session struct {
	Token      string `json:"id"`
	User_id    int    `json:"user_id"`
	Expiration string `json:"expiration"`
}
