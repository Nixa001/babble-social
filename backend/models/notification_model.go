package models

type Notification struct {
	ID               int    `json:"id"`
	Type             string `json:"type"`
	Status           bool   `json:"status"`
	UserIDSender     int    `json:"user_id_sender"`
	UserIDReceveived int    `json:"user_id_received"`
	GroupId          int    `json:"group_id"`
	Date             string `json:"date"`
}
