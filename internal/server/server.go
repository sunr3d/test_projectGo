package server

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Server *grpc.Server
	logger *zap.Logger
}

func New(logger *zap.Logger) *Server {
	return &Server{
		Server: grpc.NewServer(),
		logger: logger,
	}
}

func (s *Server) Run(port string) error {
	address := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w\n", address, err)
	}

	s.logger.Info("gRPC server start",
		zap.String("address", address),
	)

	if err = s.Server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Server) Stop() {
	s.logger.Info("gRPC server stop")
	s.Server.GracefulStop()
}
