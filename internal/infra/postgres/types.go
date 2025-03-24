package postgres_impl

import (
	"database/sql"
	"go.uber.org/zap"
)

type PostgresDB struct {
	Logger *zap.Logger
	DB     *sql.DB
}
