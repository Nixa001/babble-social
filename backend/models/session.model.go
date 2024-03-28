package models

import "time"
type Sessions struct {
	ID              int    `json:"id"`
	UserID          int    `json:"userID"`
	TokenExpiration string `json:"tokenExperation"`
}

type Session struct {
	ID         string
	UserID       int
	Expiration time.Time
}
var UserSession Session
