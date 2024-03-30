package server

import (
	db "backend/database"
	"backend/server/routes"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	Host string
	Port int
	DB   *db.Database
}

func NewServer() *Server {
	return &Server{
		Host: "localhost",
		Port: 8080,
	}
}
func (s *Server) Run() {
	fmt.Printf("Server running on port %v\nhttp://%v:%v\n", s.Port, s.Host, s.Port)
	routes := routes.Route()
	http.ListenAndServe(s.Host+":"+strconv.Itoa(s.Port), routes)
	fmt.Println("Server running on port", s.Port)
}
