package main

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omaily/autsors/config"
)

type Storage struct {
	conf *config.Storage
	pool *pgxpool.Pool
}

var pgOnce sync.Once

func NewStorage(ctx context.Context, conf *config.Storage) (*Storage, error) {
	var pgInstance *Storage
	pgOnce.Do(func() {
		dbPath := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", conf.Role, conf.Pass, conf.Host, conf.Database)
		pool, err := pgxpool.New(ctx, dbPath)
		if err != nil {
			slog.Error("unable to create connection pool", slog.String("err", err.Error()))
			return
		}
		pgInstance = &Storage{conf: conf, pool: pool}
	})
	return pgInstance, nil
}

func (s *Storage) getAmount(ctx context.Context, uuid string) (string, error) {
	query := `select name from account where id = $1`
	typeString := struct{ st string }{st: ""}

	slog.Info("Pre scan")
	row := s.pool.QueryRow(ctx, query, 1)
	err := row.Scan(&typeString.st)
	if err != nil {
		slog.Error("Error Fetching Book Details: %w", slog.String("err", err.Error()))
		return "nil", err
	}
	slog.Info("after scan")
	slog.Info(typeString.st)

	return typeString.st, nil
}

func (s *Storage) depositPay(ctx context.Context, val int) error {
	return nil
}

func (s *Storage) withdrawPay(ctx context.Context, val int) error {
	return nil
}
