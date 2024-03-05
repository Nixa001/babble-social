package server

import (
	"backend/server/routes"
	"net/http"
	"strconv"
)

type Server struct {
	Host string
	Port int
}

func NewServer() *Server {
	return &Server{
		Host: "localhost",
		Port: 8080,
	}
}

func (s *Server) Run() {
	routes := routes.Route()

	http.ListenAndServe(s.Host+":"+strconv.Itoa(s.Port), routes)

}
