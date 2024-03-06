package models

type Session struct {
	Token          string `json:"id"`
	UserId         int    `json:"user_id"`
	ExpirationDate string `json:"expiration"`
}
