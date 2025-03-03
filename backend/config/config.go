package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/chrisjpalmer/ledger/backend/internal/postgres"
	"github.com/chrisjpalmer/ledger/backend/internal/server"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"
)

// HasDotEnv - returns true if a .env file exists
func HasDotEnv(dir string) bool {
	stat, err := os.Stat(filepath.Join(dir, ".env"))
	_ = stat
	return err == nil
}

// LoadDotEnv - attempts to load the .env file
func LoadDotEnv(dir string) error {
	return godotenv.Load(filepath.Join(dir, ".env"))
}

type Config struct {
	LogLevel zapcore.Level
	Postgres postgres.Config
	Server   server.Config
}

func Load() (Config, *Errors) {
	var e Errors
	cfg := Config{
		LogLevel: parseLogLevel("APP_LOGLEVEL", &e),
		Postgres: LoadPostgresConfig(&e),
		Server:   LoadServerConfig(&e),
	}

	return cfg, &e
}

// LoadPostgresConfig - loads the config for postgres from the environment
func LoadPostgresConfig(e *Errors) postgres.Config {
	return postgres.Config{
		Database: parseString("APP_POSTGRES_DATABASE", e),
		Host:     parseString("APP_POSTGRES_HOST", e),
		Password: parseString("APP_POSTGRES_PASSWORD", e),
		Port:     parseUint16("APP_POSTGRES_PORT", e),
		User:     parseString("APP_POSTGRES_USER", e),
	}
}

// LoadServerConfig - loads the config for the server from the environment
func LoadServerConfig(e *Errors) server.Config {
	return server.Config{
		Port: parseInt("APP_SERVER_PORT", e),
	}
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
