package models

type Comment struct {
	ID      int
	Date    string
	Content string `json:"content"`
	Post_id int    `json:"post_id"`
	User_id int    `json:"user_id"`
	Media   string `json:"media"`
}

type DataComment []struct {
	Data          []Comment
	Comment_react [2]int
}
