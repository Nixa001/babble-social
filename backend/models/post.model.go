package models

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"post_title"`
	Content string `json:"post_content"`
	Media   string `json:"post_media"`
	Date    string `json:"post_date"`
	UserID  int    `json:"user_id"`
	GroupID int    `json:"group_id"`
	Type    string `json:"type"`
}
