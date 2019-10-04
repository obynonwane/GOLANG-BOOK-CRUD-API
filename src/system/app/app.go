package app

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

//Init all vals
func (s *Server) Init() {
	log.Println("Initilizing server ...")
	s.port = ":8000"
}

//Start the server
func (s *Server) Start() {
	log.Println("Starting Server")
	http.ListenAndServe(s.port, nil)
}
