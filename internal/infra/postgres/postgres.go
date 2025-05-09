package postgres_impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // Постгрес драйвер
	"go.uber.org/zap"

	"link_service/internal/config"
	"link_service/internal/interfaces/infra"
)

var _ infra.Database = (*PostgresDB)(nil)

type PostgresDB struct {
	Logger *zap.Logger
	DB     *sql.DB
}

// New Инициализация БД с проверкой соединения (конструктор).
func New(log *zap.Logger, cfg config.Postgres) (infra.Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)

	database, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err = database.Ping(); err != nil {
		return nil, fmt.Errorf("database.Ping: %w", err)
	}
	log.Info("Connect to Postgres database success")

	return &PostgresDB{Logger: log, DB: database}, nil
}

func (p *PostgresDB) Close() error {
	return p.DB.Close()
}

func (p *PostgresDB) Find(ctx context.Context, fakeLink string) (*string, error) {
	var link string
	stmt, err := p.DB.PrepareContext(ctx, "SELECT link FROM links WHERE fake_link = $1")
	if err != nil {
		return nil, fmt.Errorf("posgres.DB.PrepareContext %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, fakeLink).Scan(&link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("stmt.QueryRowContext: %w", err)
	}

	return &link, nil
}

func (p *PostgresDB) Create(ctx context.Context, link infra.InputLink) error {
	stmt, err := p.DB.PrepareContext(ctx, "INSERT INTO links (link, fake_link, erase_time) VALUES ($1,$2,$3)")
	if err != nil {
		return fmt.Errorf("p.DB.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		link.Link,
		link.FakeLink,
		link.EraseTime,
	)
	if err != nil {
		return fmt.Errorf("stmt.ExecContext: %w", err)
	}

	return nil
}
