package models

type LikePost struct {
  ID        int    `json:"id"`
  PostID    int    `json:"post_id"`
  UserID    int    `json:"user_id"`
  IsLiked   bool   `json:"isliked"`
}
