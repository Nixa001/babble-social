package models

type LikeComment struct {
	ID        int  `json:"id"`
	CommentID int  `json:"comment_id"`
	UserID    int  `json:"user_id"`
	IsLiked   bool `json:"isliked"`
}
