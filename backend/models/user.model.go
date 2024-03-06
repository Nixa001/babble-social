package models

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"user_name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserType  string `json:"user_type"`
	BirthDate string `json:"birth_date"`
	Avatar    string `json:"avatar"`
	AboutMe   string `json:"about_me"`
}
