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

//var _ postgres.Chats = (*impl)(nil) ---- Не понимаю это

type postgresDB struct {
	logger *zap.Logger
	db     *sql.DB
}

// Инициализация БД с проверкой соединения (конструктор)
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

	return &postgresDB{logger: lg, db: db}, nil
}

func (p *postgresDB) Close() error {
	return p.db.Close()
}

func (p *postgresDB) Find(ctx context.Context, fakeLink string) (string, error) {
	var link string
	stmt, err := p.db.PrepareContext(ctx, "SELECT link FROM links WHERE fake_link = ?")
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, fakeLink).Scan(&link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", status.Error(codes.Internal, err.Error())
	}

	return link, nil
}

func (p *postgresDB) Create(ctx context.Context, link infra.InputLink) (int, error) {
	id := 0
	stmt, err := p.db.PrepareContext(ctx, "INSERT INTO links (link, fake_link, erase_time) VALUES (?,?,?) RETURNING id")
	if err != nil {
		return id, status.Error(codes.Internal, err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(
		ctx,
		link.Link,
		link.FakeLink,
		link.EraseTime,
	).Scan(&id)
	if err != nil {
		return id, status.Error(codes.Internal, err.Error())
	}

	return id, nil
}
