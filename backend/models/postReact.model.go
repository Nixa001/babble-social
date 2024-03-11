package models

type PostReact struct {
	PostID   int  `json:"post_id"`
	UserID   int  `json:"user_id"`
	Reaction bool `json:"reaction"`
}
