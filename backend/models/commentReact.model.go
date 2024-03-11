package models

type CommentReact struct {
	PostID    int  `json:"id"`
	CommentID int  `json:"comment_id"`
	UserID    int  `json:"user_id"`
	Reaction  bool `json:"reaction"`
}
