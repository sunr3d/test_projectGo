package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// GetConfigFromEnv загружает конфигурации из .env.example файла и переменных окружения
func GetConfigFromEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Не удалось загрузить .env файл: %s\n", err.Error())
	}

	// Инициализация структуры конфигурации
	cfg := &Config{}
	// Парсинг переменных окружения в структуру
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
