package models

type Viewers []struct {
	Post_id int `json:"postId"`
	User_id int `json:"userId"`
}
