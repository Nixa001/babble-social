package models

type UserFollowers struct {
	ID        int    `json:"id"`
	UserFollowed    int    `json:"user_followed"`
	UserFollower   int    `json:"user_follower"`
}
