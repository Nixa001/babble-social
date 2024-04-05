package service

var (
	AuthServ    AuthService
	PostServ    PostService
	CommentServ CommentService
	FollowServ  FollowService
)

func init() {
	AuthServ.init()
	PostServ.init()
	CommentServ.init()
	FollowServ.init()
}
