package postgres

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"link_service/internal/interfaces/infra"
)

//var _ postgres.Chats = (*impl)(nil) ---- Не понимаю это

type PostgresDB struct {
	db *sql.DB
}

// Инициализация БД с проверкой соединения (конструктор)
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

func (p *PostgresDB) Find(ctx context.Context, fakeLink string) (string, error) {
	var link string
	stmt, err := p.db.Prepare("SELECT link FROM links WHERE fake_link = ?")
	if err != nil {
		return "", status.Error(codes.Internal, err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, fakeLink).Scan(&link)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", status.Errorf(codes.NotFound, "link %s not found", fakeLink)
		}
		return "", status.Error(codes.Internal, err.Error())
	}

	return link, nil
}

func (p *PostgresDB) Create(ctx context.Context, link infra.InputLink) error {
	stmt, err := p.db.Prepare("INSERT INTO links (link, fake_link, erase_time) VALUES (?,?,?)")
	if err != nil {
		return status.Error(codes.Internal, err.Error())
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
