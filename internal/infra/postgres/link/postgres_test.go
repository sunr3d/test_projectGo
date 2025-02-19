package postgres_impl_test

import (
	"context"
	"database/sql"
	"errors"
	"link_service/internal/interfaces/infra"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"link_service/internal/infra/postgres/link"
)

func TestPostgresDB_Find(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop()
	postgresDB := &postgres_impl.PostgresDB{Logger: logger, Db: db}

	fakeLink := "http://fake.com"
	expectedLink := "http://example.com"

	rows := sqlmock.NewRows([]string{"link"}).AddRow(expectedLink)
	mock.ExpectPrepare("SELECT link FROM links WHERE fake_link = \\$1").
		ExpectQuery().
		WithArgs(fakeLink).
		WillReturnRows(rows)

	link, err := postgresDB.Find(context.Background(), fakeLink)
	assert.NoError(t, err)
	assert.Equal(t, expectedLink, *link)
}

func TestPostgresDB_Find_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop()
	postgresDB := &postgres_impl.PostgresDB{Logger: logger, Db: db}

	fakeLink := "http://nonexistent.com"

	mock.ExpectPrepare("SELECT link FROM links WHERE fake_link = \\$1").
		ExpectQuery().
		WithArgs(fakeLink).
		WillReturnError(sql.ErrNoRows)

	link, err := postgresDB.Find(context.Background(), fakeLink)
	assert.NoError(t, err)
	assert.Nil(t, link)
}

func TestPostgresDB_Find_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop()
	postgresDB := &postgres_impl.PostgresDB{Logger: logger, Db: db}

	fakeLink := "http://error.com"

	mock.ExpectPrepare("SELECT link FROM links WHERE fake_link = \\$1").
		ExpectQuery().
		WithArgs(fakeLink).
		WillReturnError(errors.New("query error"))

	link, err := postgresDB.Find(context.Background(), fakeLink)
	assert.Error(t, err)
	assert.Nil(t, link)
}

func TestPostgresDB_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop()
	postgresDB := &postgres_impl.PostgresDB{Logger: logger, Db: db}

	inputLink := infra.InputLink{
		Link:      "http://example.com",
		FakeLink:  "http://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	mock.ExpectPrepare("INSERT INTO links \\(link, fake_link, erase_time\\) VALUES \\(\\$1,\\$2,\\$3\\)").
		ExpectExec().
		WithArgs(inputLink.Link, inputLink.FakeLink, inputLink.EraseTime).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = postgresDB.Create(context.Background(), inputLink)
	assert.NoError(t, err)
}

func TestPostgresDB_Create_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	logger := zap.NewNop()
	postgresDB := &postgres_impl.PostgresDB{Logger: logger, Db: db}

	inputLink := infra.InputLink{
		Link:      "http://example.com",
		FakeLink:  "http://fake.com",
		EraseTime: time.Now().Add(24 * time.Hour),
	}

	mock.ExpectPrepare("INSERT INTO links \\(link, fake_link, erase_time\\) VALUES \\(\\$1,\\$2,\\$3)").
		ExpectQuery().
		WithArgs(inputLink.Link, inputLink.FakeLink, inputLink.EraseTime).
		WillReturnError(errors.New("insert error"))

	err = postgresDB.Create(context.Background(), inputLink)
	assert.Error(t, err)
}
