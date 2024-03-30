package models

type Event struct {
	ID          int    `json:"id"`
	GroupID     int    `json:"group_id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
