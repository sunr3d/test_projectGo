package postgres

import (
	"database/sql"
)

type PostgresDB struct {
	db *sql.DB
}

// Инициализация БД с проверкой соединения
func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}
