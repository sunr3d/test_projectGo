package entrypoint

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap" // logger
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection" // reflection для теста ручек через `grpcurl` в терминале

	"link_service/internal/config"
	hh "link_service/internal/handlers/health"
	lsh "link_service/internal/handlers/link_service"
	kafka_impl "link_service/internal/infra/kafka"
	postgres_impl "link_service/internal/infra/postgres"
	redis_impl "link_service/internal/infra/redis"
	"link_service/internal/server"
	"link_service/internal/service/link_service_impl"
	pbls "link_service/proto/link_service"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	/// Слой репозиториев (infra)
	// Коннект к БД Постгрес по данным из конфига
	postgresRepo, err := postgres_impl.New(logger, cfg.Postgres)
	if err != nil {
		return fmt.Errorf("create postgres link service: %w", err)
	}

	// Коннект к Редису (как кэш БД) по данным из конфига
	redisRepo, err := redis_impl.New(logger, cfg.Redis)
	if err != nil {
		return fmt.Errorf("create redis link service: %w", err)
	}

	// Коннект к Кафке по данным из конфига
	kafkaWriter, err := kafka_impl.New(logger, cfg.KafkaPort)
	if err != nil {
		return fmt.Errorf("create kafka link service: %w", err)
	}

	/// Сервисный слой
	svc := link_service_impl.New(logger, postgresRepo, redisRepo, kafkaWriter)

	grpcServer := server.New(logger, cfg)
	reflection.Register(grpcServer.Server) // reflection для теста ручек через `grpcurl` в терминале

	// Регистрация сервиса health
	healthService := hh.New()
	grpc_health_v1.RegisterHealthServer(grpcServer.Server, healthService)

	// Регистрация сервиса link_service
	linkService := lsh.New(svc)
	pbls.RegisterLinkServiceServer(grpcServer.Server, linkService)

	/// Запуск сервера
	go func() {
		if err = grpcServer.Run(); err != nil {
			logger.Fatal("grpc server run error", zap.Error(err))
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	<-done
	grpcServer.Stop()

	return nil

}
