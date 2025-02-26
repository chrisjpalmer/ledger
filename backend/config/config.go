package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chrisjpalmer/ledger/backend/internal/postgres"
	"github.com/chrisjpalmer/ledger/backend/internal/server"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel zapcore.Level
	Postgres postgres.Config
	Server   server.Config
}

func Load() (Config, *Errors) {
	var e Errors
	cfg := Config{
		LogLevel: parseLogLevel("APP_LOGLEVEL", &e),
		Postgres: postgres.Config{
			Database: parseString("APP_POSTGRES_DATABASE", &e),
			Host:     parseString("APP_POSTGRES_HOST", &e),
			Password: parseString("APP_POSTGRES_PASSWORD", &e),
			Port:     parseUint16("APP_POSTGRES_PORT", &e),
			User:     parseString("APP_POSTGRES_USER", &e),
		},
		Server: server.Config{
			Port: parseInt("APP_SERVER_PORT", &e),
		},
	}

	return cfg, &e
}

func parseString(key string, e *Errors) string {
	val, ok := mustEnv(key, e)
	if !ok {
		return ""
	}

	return val
}

func parseUint16(key string, e *Errors) uint16 {
	val, ok := mustEnv(key, e)
	if !ok {
		return 0
	}

	it, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		e.Add(err)
	}

	return uint16(it)
}

func parseInt(key string, e *Errors) int {
	val, ok := mustEnv(key, e)
	if !ok {
		return 0
	}

	it, err := strconv.Atoi(val)
	if err != nil {
		e.Add(err)
	}

	return it
}

func parseLogLevel(key string, e *Errors) zapcore.Level {
	val, ok := mustEnv(key, e)
	if !ok {
		return zapcore.InfoLevel
	}

	lvl, err := zapcore.ParseLevel(val)
	if err != nil {
		e.Add(err)
		return zapcore.InfoLevel
	}

	return lvl
}

func mustEnv(key string, e *Errors) (string, bool) {
	val := os.Getenv(key)
	if val == "" {
		e.Add(fmt.Errorf("env var %s was empty", key))
		return "", false
	}
	return val, true
}
