package main

import (
	"log"

	"go.uber.org/zap"

	"link_service/internal/config"
	"link_service/internal/entrypoint"
	"link_service/internal/logger"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("config.GetConfigFromEnv: %s\n", err.Error())
	}

	// Инициализация логгера
	zapLogger := logger.NewClientZapLogger(cfg.LogLevel, cfg.ServiceName)

	// Запуск сервера
	if err = entrypoint.Run(cfg, zapLogger); err != nil {
		zapLogger.Fatal("entrypoint.Run: ", zap.Error(err))
	}
}
