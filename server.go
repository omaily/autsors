package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/omaily/autsors/config"
)

type APIServer struct {
	conf    *config.HTTPServer
	storage *Storage
}

func NewServer(ctx context.Context, conf *config.Config) (*APIServer, error) {

	http := &conf.HTTPServer
	if http == nil {
		return nil, errors.New("configuration files are not initialized")
	}
	if http.Address == "" || http.Port == "" {
		return nil, errors.New("configuration address cannot be blank")
	}

	storage, err := NewStorage(ctx, &conf.Storage)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not initialize storage: %s", err))
	}

	return &APIServer{
		conf:    http,
		storage: storage,
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
	defer s.storage.pool.Close()
	return srv.Shutdown(ctx)
}

func (s *APIServer) router() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("POST /api/v1/wallet", postWallet(s.storage))
	router.HandleFunc("GET /api/v1/wallets/{uuid}", getWallet(s.storage))

	return router
}
