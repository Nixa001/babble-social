package models

type Sessions struct {
	ID              int    `json:"id"`
	UserID          int    `json:"userID"`
	TokenExpiration string `json:"tokenExperation"`
}
