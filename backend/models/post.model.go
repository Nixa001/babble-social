package models

type Post struct {
	ToIns      InsPost  `json: "insert"`
	Categories Category `json: "categories"`
	Viewers    Viewers  `json:"viewers"`
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

type Category []struct {
	Post_id  int    `json:"postId"`
	Category string `json:"category"`
}

type Viewers []struct {
	Post_id int `json:"postId"`
	User_id int `json:"userId"`
}

type DataPost struct {
	ID         int
	Date       string
	Avatar     string
	Content    string
	Categories string
	Media      string
	User_id    int
	Group_id   int
	Privacy    string
	FullName   string
	UserName   string
	comments   int
}
