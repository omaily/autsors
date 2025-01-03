package main

import (
	"context"
	"errors"
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
}

var pgOnce sync.Once

func NewStorage(ctx context.Context, conf *config.Storage) (*Storage, error) {
	var pgInstance *Storage
	pgOnce.Do(func() {
		dbPath := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", conf.Role, conf.Pass, conf.Host, conf.Port, conf.Database)
		pool, err := pgxpool.New(ctx, dbPath)
		if err != nil {
			slog.Error("unable to create connection pool", slog.String("err", err.Error()))
			return
		}
		pgInstance = &Storage{conf: conf, pool: pool}
	})
	return pgInstance, nil
}

func (s *Storage) getAmount(ctx context.Context, uuid string) (int, error) {
	query := `select amount from account where uuid = $1`
	var cash int
	row := s.pool.QueryRow(ctx, query, uuid)
	err := row.Scan(&cash)
	if err != nil {
		slog.Error("Error Fetching Book Details: %w", slog.String("err", err.Error()))
		return 0, err
	}

	return cash, nil
}

func (s *Storage) depositPay(ctx context.Context, uuid string, amount int) error {
	cash, err := s.getAmount(ctx, uuid)
	if err != nil {
		return err
	}

	cash += amount

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	args := pgx.NamedArgs{
		"cash": cash,
		"uuid": uuid,
	}

	_, err = tx.Exec(ctx, "UPDATE account SET amount = @cash where uuid = @uuid", args)
	if err != nil {
		slog.Error("UPDATE fall: %w", slog.String("err", err.Error()))
		return err
	}

	slog.Info("user depositPay", slog.String("user", uuid))
	return nil
}

func (s *Storage) withdrawPay(ctx context.Context, uuid string, amount int) error {

	cash, err := s.getAmount(ctx, uuid)
	if err != nil {
		return err
	}

	if amount > cash {
		return errors.New("there are insufficient funds in your account")
	}

	cash -= amount

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	args := pgx.NamedArgs{
		"cash": cash,
		"uuid": uuid,
	}

	_, err = tx.Exec(ctx, "UPDATE account SET amount = @cash where uuid = @uuid", args)
	if err != nil {
		slog.Error("UPDATE fall: %w", slog.String("err", err.Error()))
		return err
	}

	slog.Info("user withdrawPay", slog.String("user", uuid))
	return nil
}
