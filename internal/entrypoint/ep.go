package entrypoint

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"link_service/internal/config"
	"net"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	grpcServer := grpc.NewServer()

	// TODO: Сервисы, первый должен быть HealthCheck

	address := fmt.Sprintf(":%s", cfg.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen on address %s: %w\n", address, err)
	}

	logger.Info("start gRPC server",
		zap.String("port", cfg.GRPCPort),
		zap.String("version", cfg.Version),
	)

	if err = grpcServer.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil

}
