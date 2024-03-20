package models

type Post struct {
	ToIns      InsPost  `json: "insert"`
	Categories []string `json: "categories"`
	Viewers    string  `json:"viewers"`
}
type InsPost struct {
	ID       int
	Date     string
	Content  string `json:"content"`
	Media    string `json:"media"`
	User_id  int    `json:"userID"`
	Group_id int    `json:"groupID"`
	Privacy  string `json:"privacy"`
}

type Category struct {
	Post_id  string    `json:"postId"`
	Category string `json:"category"`
}

type DataPost struct {
	Avatar     string
	Categories string
	comments   int
	Content    string
	Date       string
	FullName   string
	Group_id   int
	ID         int
	Media      string
	Privacy    string
	UserName   string
	User_id    int
	Viewers    string
}
