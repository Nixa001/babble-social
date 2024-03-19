package models

type Comment struct {
	ID      int
	Date    string
	Content string `json:"content"`
	Post_id int    `json:"postID"`
	User_id int    `json:"userID"`
	Media   string `json:"media"`
}

type DataComment []struct {
	Avatar   string
	Content  string
	Date     string
	FullName string
	ID       int
	Media    string
	Post_id  string
	Username string
}
