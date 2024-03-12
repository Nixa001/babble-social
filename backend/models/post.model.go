package models

type Post struct {
	ID       int
	Date     string
	Content  string
	Media    string
	User_id  int
	Group_id int
	Privacy  string
}

type DataPost []struct {
	Data          []Post
	comments      DataComment
	Post_reaction [2]int
}
