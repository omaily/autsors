package main

import (
	"context"
	"errors"
	"fmt"
	"june+/config"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pgOnce sync.Once

type APIServer struct {
	conf    *config.HTTPServer
	storage *pgxpool.Pool
}

func NewServer(ctx context.Context, conf *config.Config) (*APIServer, error) {

	http := &conf.HTTPServer
	if http == nil {
		return nil, errors.New("configuration files are not initialized")
	}
	if http.Address == "" || http.Port == "" {
		return nil, errors.New("configuration address cannot be blank")
	}

	var pgInstance *pgxpool.Pool
	pgOnce.Do(func() {
		pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s/%s", (*conf).Role, (*conf).Pass, (*conf).Host /*,(*cs).Port*/, (*conf).Database))
		if err != nil {
			slog.Error("unable to create connection pool: %w", err)
			return
		}
		pgInstance = pool
	})

	return &APIServer{
		conf:    http,
		storage: pgInstance,
	}, nil
}

func (s *APIServer) Start() error {

	srv := &http.Server{
		Addr:         s.conf.Port,
		Handler:      s.router(),
		ReadTimeout:  s.conf.Timeout * time.Second,
		WriteTimeout: s.conf.Timeout * 2 * time.Second,
		IdleTimeout:  s.conf.IdleTimeout * time.Second,
	}

	go func() {
		slog.Info("starting server", slog.String("addres", s.conf.Address))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("не маслает", slog.String("err", err.Error()))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	slog.Info("stopping server due to syscall or collapse")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return srv.Shutdown(ctx)
}

func (s *APIServer) router() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("POST /api/v1/wallet", postWallet)
	router.HandleFunc("GET /api/v1/wallet/{uuid}", getWallet)

	return router
}
