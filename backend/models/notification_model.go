package models

type Notification struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Type             string `json:"type"`
	UserIDSender     int    `json:"user_id_sender"`
	UserIDReceveived int    `json:"user_id_received"`
	Date             string `json:"date"`
	Status           bool   `json:"status"`
}
