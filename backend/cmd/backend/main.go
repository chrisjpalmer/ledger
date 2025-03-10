package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/chrisjpalmer/ledger/backend/config"
	"github.com/chrisjpalmer/ledger/backend/internal/postgres"
	"github.com/chrisjpalmer/ledger/backend/internal/server"
	"go.uber.org/zap"
)

func main() {
	// create logger
	zl, lvl, err := newLogger()
	if err != nil {
		log.Printf("unable to create logger %s", err.Error())
	}

	// load env vars
	if config.HasDotEnv(".") {
		zl.Info("detected .env file... loading it in")
		if err := config.LoadDotEnv("."); err != nil {
			zl.Fatal("unable to load environment vars: %w", zap.Error(err))
		}
	}

	// load config
	cfg, e := config.Load()
	if e.HasErrors() {
		e.ForEach(func(err error) { zl.Error("env var parsing error", zap.Error(err)) })
		zl.Fatal("found errors while parsing the environment variables")
	}

	// set log level
	zl.Info("setting log level", zap.String("new_log_level", cfg.LogLevel.String()))
	lvl.SetLevel(cfg.LogLevel)

	// postgres
	p, err := postgres.New(zl, cfg.Postgres)
	if err != nil {
		zl.Fatal("error while connecting to postgres", zap.Error(err))
	}

	// start server
	lisErr := make(chan error, 1)
	srv := server.NewServer(zl, p, cfg.Server)
	go func() {
		if err := srv.Listen(); err != nil {
			lisErr <- err
		}
	}()

	// wait for ctrl c
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	select {
	case <-s:
		zl.Info("received shutdown signal...")
	case <-lisErr:
		zl.Error("server closed prematurely with error", zap.Error(err))
	}

	// close
	srv.Close()
	p.Close()
}

func newLogger() (*zap.Logger, zap.AtomicLevel, error) {
	cfg := zap.NewProductionConfig()

	zl, err := cfg.Build()
	if err != nil {
		return nil, zap.AtomicLevel{}, err
	}

	return zl, cfg.Level, nil

}
