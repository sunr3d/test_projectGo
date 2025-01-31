package entrypoint

import (
	"fmt"
	"go.uber.org/zap" // logger
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection" // reflection для теста ручек через `grpcurl` в терминале
	"link_service/internal/config"
	"link_service/internal/handlers/health_handler"
	"net"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer) // reflection для теста ручек через `grpcurl` в терминале

	// Регистрация сервиса health
	healthService := health_handler.NewHealthService()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthService)

	// Само поднятие сервера
	address := fmt.Sprintf(":%s", cfg.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w\n", address, err)
	}

	logger.Info("gRPC server start",
		zap.String("address", "localhost:"),
		zap.String("port", cfg.GRPCPort),
		zap.String("version", cfg.Version),
	)

	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil

}
