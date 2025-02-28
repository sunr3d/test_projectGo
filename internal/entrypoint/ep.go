package entrypoint

import (
	"fmt"

	"go.uber.org/zap" // logger
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection" // reflection для теста ручек через `grpcurl` в терминале

	"link_service/internal/config"
	hh "link_service/internal/handlers/health"
	lsh "link_service/internal/handlers/link_service"
	postgres_impl "link_service/internal/infra/postgres/link"
	redis_impl "link_service/internal/infra/redis"
	"link_service/internal/server"
	"link_service/internal/service/link_service_impl"
	pbls "link_service/proto/link_service"
)

func Run(cfg *config.Config, logger *zap.Logger) error {

	/// Слой репозиториев (infra)
	// Коннект к БД Постгрес по данным из конфига
	pg, err := postgres_impl.New(logger, cfg.Postgres)
	if err != nil {
		return fmt.Errorf("create postgres link service: %w", err)
	}

	// Коннект к Редису (как кэш БД) по данным из конфига
	rd, err := redis_impl.New(logger, cfg.Redis)
	if err != nil {
		return fmt.Errorf("create redis link service: %w", err)
	}

	/// Сервисный слой
	svc := link_service_impl.New(logger, pg, rd)

	grpcServer := server.New(logger)
	reflection.Register(grpcServer.Server) // reflection для теста ручек через `grpcurl` в терминале

	// Регистрация сервиса health
	healthService := hh.New()
	grpc_health_v1.RegisterHealthServer(grpcServer.Server, healthService)

	// Регистрация сервиса link_service
	linkService := lsh.New(svc)
	pbls.RegisterLinkServiceServer(grpcServer.Server, linkService)

	/// Запуск сервера
	if err = grpcServer.Run(cfg.GRPCPort); err != nil {
		return fmt.Errorf("run grpc server: %w", err)
	}

	return nil

}
