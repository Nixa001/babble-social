package models

type Post struct {
	ID       int
	Date     string
	Content  string   `json:"content"`
	Media    string   `json:"media"`
	User_id  int      `json:"userID"`
	Group_id int      `json:"groupID"`
	Privacy  string   `json:"privacy"`
	Viewers  []string `json:"viewers"`
}

type DataPost []struct {
	Avatar        string
	Data          []Post
	FullName      string
	Username      string
	comments      DataComment
	Post_reaction [2]int
}
