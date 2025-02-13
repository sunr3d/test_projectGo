package entrypoint

import (
	"go.uber.org/zap" // logger
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection" // reflection для теста ручек через `grpcurl` в терминале
	"link_service/internal/config"
	hh "link_service/internal/handlers/health"
	lsh "link_service/internal/handlers/link_service"
	postgres_impl "link_service/internal/infra/postgres/link"
	"link_service/internal/server"
	"link_service/internal/service/link_service_impl"
	pbls "link_service/proto/link_service"
)

func Run(cfg *config.Config, logger *zap.Logger) error {

	//infra
	//TODO change
	pg, err := postgres_impl.New(logger, cfg.Postgres)
	if err != nil {
		return err
	}

	//service layer
	svc := link_service_impl.New(logger, pg)

	grpcServer := server.NewServer(logger)
	reflection.Register(grpcServer.Server) // reflection для теста ручек через `grpcurl` в терминале

	// Регистрация сервиса health
	healthService := hh.NewHealthService()
	grpc_health_v1.RegisterHealthServer(grpcServer.Server, healthService)

	// Регистрация сервиса link_service
	linkService := lsh.NewLinkService(svc) // placeholder
	pbls.RegisterLinkServiceServer(grpcServer.Server, linkService)

	// Запуск сервераЗ
	if err := grpcServer.Run(cfg.GRPCPort); err != nil {
		return err
	}

	return nil

}
