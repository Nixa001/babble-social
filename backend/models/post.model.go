package models

type Post struct {
	ID            int
	Date          string
	Comments      Comments
	Post_reaction [2]int
	Content       string `json:"post_content"`
	Media         string `json:"post_media"`
	UserID        int    `json:"user_id"`
	GroupID       int    `json:"group_id"`
	Type          string `json:"type"`
}
type Posts []Post
