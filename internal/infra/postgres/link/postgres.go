package postgres_impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"link_service/internal/config"
	"link_service/internal/interfaces/infra"
)

var _ infra.Database = (*PostgresDB)(nil)

type PostgresDB struct {
	Logger *zap.Logger
	Db     *sql.DB
}

// New Инициализация БД с проверкой соединения (конструктор)
func New(lg *zap.Logger, cfg config.Postgres) (infra.Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err = db.Ping(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	lg.Info("Connect to Postgres database success")

	return &PostgresDB{Logger: lg, Db: db}, nil
}

func (p *PostgresDB) Close() error {
	return p.Db.Close()
}

func (p *PostgresDB) Find(ctx context.Context, fakeLink string) (*string, error) {
	var link string
	stmt, err := p.Db.PrepareContext(ctx, "SELECT link FROM links WHERE fake_link = $1")
	if err != nil {
		return nil, fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, fakeLink).Scan(&link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &link, nil
}

func (p *PostgresDB) Create(ctx context.Context, link infra.InputLink) error {
	stmt, err := p.Db.PrepareContext(ctx, "INSERT INTO links (link, fake_link, erase_time) VALUES ($1,$2,$3)")
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		link.Link,
		link.FakeLink,
		link.EraseTime,
	)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
