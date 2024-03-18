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
	FullName      string
	Username      string
	Avatar        string
	Data          []Comment
	Comment_react [2]int
}
