package server

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"link_service/internal/config"
	"link_service/internal/gateway"
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

func (s *Server) Run(ctx context.Context, cfg *config.Config) error {
	address := fmt.Sprintf(":%s", cfg.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w\n", address, err)
	}

	s.logger.Info("gRPC server start",
		zap.String("address", address),
	)

	if cfg.GatewayFlag {
		go func() {
			gw := gateway.New(s.logger)
			if err = gw.Run(ctx, cfg.GRPCPort, cfg.HTTPPort); err != nil {
				s.logger.Error("server.Run: ", zap.Error(err))
			}
		}()
	}

	if err = s.Server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}

func (s *Server) Stop() {
	s.logger.Info("gRPC server stop")
	s.Server.GracefulStop()
}
