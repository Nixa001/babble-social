package models

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"comment_content"`
	Date    string `json:"comment_date"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Type    string `json:"type"`
	Media   string `json:"comment_media"`
	IsLiked bool   `json:"isliked"`
}
