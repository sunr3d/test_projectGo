package entrypoint

import (
	"go.uber.org/zap" // logger
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection" // reflection для теста ручек через `grpcurl` в терминале
	"link_service/internal/config"
	"link_service/internal/handlers/health_handler"
	"link_service/internal/handlers/link_service"
	"link_service/internal/server"
	pbls "link_service/proto/link_service"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	grpcServer := server.NewServer(logger)
	reflection.Register(grpcServer.Server) // reflection для теста ручек через `grpcurl` в терминале

	// Регистрация сервиса health
	healthService := health_handler.NewHealthService()
	grpc_health_v1.RegisterHealthServer(grpcServer.Server, healthService)

	// Регистрация сервиса link_service
	linkService := link_service.NewLinkService(nil) // placeholder
	pbls.RegisterLinkServiceServer(grpcServer.Server, linkService)

	// Запуск сервераЗ
	if err := grpcServer.Run(cfg.GRPCPort); err != nil {
		return err
	}

	return nil

}
