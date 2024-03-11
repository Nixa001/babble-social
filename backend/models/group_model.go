package models

type Group struct {
	ID           int    `json:"id"`
	Name         string `json:"group_name"`
	Description  string `json:"group_description"`
	UserIDOwner  int    `json:"user_id_owner"`
	DateCreation string `json:"dateCreation"`
	Avatar       string `json:"avatar"`
}
