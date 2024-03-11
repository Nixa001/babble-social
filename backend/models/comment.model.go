package models

type Comment struct {
	ID            int
	Date          string
	Comment_react [2]int
	Content       string `json:"comment_content"`
	PostID        int    `json:"post_id"`
	UserID        int    `json:"user_id"`
	Media         string `json:"comment_media"`
}

type Comments []Comment
