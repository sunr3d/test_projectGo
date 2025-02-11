package postgres

import (
	"database/sql"
	"link_service/internal/interfaces/infra"
)

//var _ postgres.Chats = (*impl)(nil)

type PostgresDB struct {
	db *sql.DB
}

// Инициализация БД с проверкой соединения
func NewPostgresDB(dsn string) (infra.Database, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}
