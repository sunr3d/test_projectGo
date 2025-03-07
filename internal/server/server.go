package server

import (
	"context"
	"fmt"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"link_service/internal/server/gateway"
)

type Server struct {
	Server        *grpc.Server
	logger        *zap.Logger
	GRPCAddress   string
	HTTPAddress   string
	GatewayEnable bool
	ctx           context.Context
	cancel        context.CancelFunc
}

func New(logger *zap.Logger, GRPCPort, HTTPPort string, GatewayEnable bool) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		Server:        grpc.NewServer(),
		logger:        logger,
		GRPCAddress:   fmt.Sprintf("localhost:%s", GRPCPort),
		HTTPAddress:   fmt.Sprintf("localhost:%s", HTTPPort),
		GatewayEnable: GatewayEnable,
		ctx:           ctx,
		cancel:        cancel,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.GRPCAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w\n", s.GRPCAddress, err)
	}

	s.logger.Info("gRPC server start",
		zap.String("address", s.GRPCAddress),
	)

	if s.GatewayEnable {
		go func() {
			gw := gateway.New(s.logger)
			if err = gw.Run(s.ctx, s.GRPCAddress, s.HTTPAddress); err != nil {
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
	s.cancel()
	s.Server.GracefulStop()
}
