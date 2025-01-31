package entrypoint

import (
	"fmt"
	"go.uber.org/zap" // logger
	"google.golang.org/grpc"
	"link_service/internal/config"
	//"link_service/internal/handlers"
	"net"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	grpcServer := grpc.NewServer()

	// TODO: Регистрация сервисов, первый должен быть HealthCheck

	address := fmt.Sprintf(":%s", cfg.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w\n", address, err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}
	logger.Info("gRPC server started successfully",
		zap.String("address", "localhost:"),
		zap.String("port", cfg.GRPCPort),
		zap.String("version", cfg.Version),
	)
	return nil

}
