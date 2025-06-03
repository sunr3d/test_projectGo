# test_projectGo — Микросервис для редиректа ссылок

Пет-проект для изучения микросервисной архитектуры на Go с использованием gRPC, PostgreSQL, Redis, Kafka, ClickHouse и Docker Compose.

## Описание

Сервис реализует систему редиректа ссылок с постоянным хранением, кешированием, метриками, авторизацией на основе JWT и интеграцией с Kafka для стриминга данных.

## Функционал

- gRPC-сервер с ручкой health check
- Мини-сервис редиректов:
    - POST Add: добавление ссылки (link, fakeLink, erase time) в PostgreSQL
    - GET Take: получение ссылки по URI
- Кеширование в Redis для быстрого доступа
- gRPC Gateway для поддержки HTTP/JSON
- Метрики Prometheus на ручках
- Ручка-продюсер в Kafka и консюмер, пишущий в ClickHouse

## Структура проекта

- test_projectGo/ — основной сервис и исходный код
    - api/ — protobuf-описания для gRPC и gRPC Gateway
    - cmd/ — точки входа для запуска сервиса
    - config/ — конфигурационные файлы и структуры
    - internal/ — бизнес-логика (handlers, services, repositories и др.)
    - migrations/ — SQL-миграции для PostgreSQL
    - proto/ — сгенерированный Go-код из protobuf

## Технологии

- Go (Golang)
- gRPC и gRPC Gateway
- PostgreSQL
- Redis
- Kafka
- ClickHouse
- Docker и Docker Compose
- Prometheus
- Goose (миграции БД)

## Быстрый старт

1. Клонируйте репозиторий.
2. Скопируйте .env.example в .env и при необходимости измените настройки.
3. Запустите все сервисы:
      docker-compose up -d
   
4. Примените миграции к базе данных:
      task migrate-up
   
5. Соберите и запустите Go-сервис:
      task container-build
   
6. Kafka UI будет доступен по адресу http://localhost:8080, Prometheus — по адресу http://localhost:9090.
