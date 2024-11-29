package main

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaily/autsors/config"
)

type Storage struct {
	conf *config.Storage
	pool *pgxpool.Pool
	User
}

var pgOnce sync.Once

func NewStorage(ctx context.Context, conf *config.Storage) (*Storage, error) {
	var pgInstance *Storage
	pgOnce.Do(func() {
		dbPath := fmt.Sprintf("postgres://%s:%s@%s/%s", conf.Role, conf.Pass, conf.Host /*,(*cs).Port*/, conf.Database)
		pool, err := pgxpool.New(ctx, dbPath)
		if err != nil {
			slog.Error("unable to create connection pool: %w", err)
			return
		}
		pgInstance = &Storage{conf: conf, pool: pool}
	})
	return pgInstance, nil
}

func (s *Storage) getAmount(ctx context.Context, uuid string) (int, error) {
	query := `select amount from account where uuid = @uuid`
	args := pgx.NamedArgs{
		"uuid": uuid,
	}
	row := s.pool.QueryRow(context.Background(), query, args)

	var val int
	err := row.Scan(&val)
	if err != nil {
		slog.Error("Error Fetching Book Details: %w", err)
		return 0, err
	}

	return val, nil
}

func (s *Storage) depositPay(ctx context.Context, val int) error {
	return nil
}

func (s *Storage) withdrawPay(ctx context.Context, val int) error {
	return nil
}
