package models

import "github.com/gofrs/uuid/v5"

type Session struct {
	Id             uuid.UUID `json:"id"`
	ExpirationDate string    `json:"expirationDate"`
	Data           string    `json:"data"`
}
