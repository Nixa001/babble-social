package models

type WResponse struct {
	Type       string
	Data       any
	StatusCode int
	Display    bool
	Msg        string
}
