package server

import "google.golang.org/grpc"

type Server struct {
	server *grpc.Server
}

func NewServer() *Server {
	return &Server{
		server: grpc.NewServer(),
	}
}

func (s *Server) Run() {

}

func (s *Server) Stop() {}
