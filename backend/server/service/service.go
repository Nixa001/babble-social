package service

var (
	AuthServ AuthService
	PostServ PostService
)

func init() {
	AuthServ.init()
	PostServ.init()
}
