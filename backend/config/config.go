package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chrisjpalmer/ledger/backend/internal/server"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Server   server.Config
	LogLevel zapcore.Level
}

func Load() (Config, *Errors) {
	var e Errors
	cfg := Config{
		LogLevel: parseLogLevel("APP_LOGLEVEL", &e),
		Server: server.Config{
			Port: parseInt("APP_SERVER_PORT", &e),
		},
	}

	return cfg, &e
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
