package entrypoint

import (
	"go.uber.org/zap" // logger
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection" // reflection для теста ручек через `grpcurl` в терминале
	"link_service/internal/config"
	"link_service/internal/handlers/health_handler"
	"link_service/internal/server"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	grpcServer := server.NewServer(logger)
	reflection.Register(grpcServer.Server) // reflection для теста ручек через `grpcurl` в терминале

	// Регистрация сервиса health
	healthService := health_handler.NewHealthService()
	grpc_health_v1.RegisterHealthServer(grpcServer.Server, healthService)

	// Запуск сервераЗ
	if err := grpcServer.Run(cfg.GRPCPort); err != nil {
		return err
	}

	return nil

}
