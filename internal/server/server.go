package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"link_service/internal/config"
	"link_service/internal/interceptors"
	"link_service/internal/server/gateway"
	"link_service/internal/server/metrics"
)

type Server struct {
	GRPCAddress    string
	HTTPAddress    string
	PrometheusAddr string
	Server         *grpc.Server
	logger         *zap.Logger
	ctx            context.Context
	cancel         context.CancelFunc
	GatewayEnable  bool
}

func New(logger *zap.Logger, cfg *config.Config) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	mtrInterceptor := interceptors.MetricsUnaryInterceptor()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(mtrInterceptor))

	return &Server{
		GRPCAddress:    fmt.Sprintf("localhost:%s", cfg.GRPCPort),
		HTTPAddress:    fmt.Sprintf("localhost:%s", cfg.HTTPPort),
		PrometheusAddr: fmt.Sprintf("localhost:%s", cfg.PrometheusPort),
		Server:         grpcServer,
		logger:         logger,
		ctx:            ctx,
		cancel:         cancel,
		GatewayEnable:  cfg.GatewayEnable,
	}
}

func (s *Server) Run() error {
	listener, err := net.Listen("tcp", s.GRPCAddress)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w\n", s.GRPCAddress, err)
	}

	s.logger.Info("server.Run: gRPC server started",
		zap.String("address", s.GRPCAddress),
	)

	go func() {
		mtr := metrics.New(s.logger)
		if err = mtr.Init(); err != nil {
			s.logger.Error("metrics.Init: ", zap.Error(err))
			return
		}

		if err = mtr.Run(s.ctx, s.PrometheusAddr); err != nil {
			s.logger.Error("metrics.Run: ", zap.Error(err))
		}
	}()

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
	s.logger.Info("gRPC server STOP signal")
	s.cancel()
	time.Sleep(500 * time.Millisecond)
	s.Server.GracefulStop()
	s.logger.Info("gRPC server stopped")
}
