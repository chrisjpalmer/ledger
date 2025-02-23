package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Postgres struct {
	pool *pgxpool.Pool
	zl   *zap.Logger
}

type Config struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}

func New(zl *zap.Logger, cfg Config) (*Postgres, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	c, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("error while parsing config: %w", err)
	}

	c.MinConns = 100
	c.MaxConns = 1000
	c.MaxConnLifetime = time.Minute * 5
	c.MaxConnIdleTime = time.Minute * 1

	pool, err := pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("error while creating pgx pool %w", err)
	}

	return &Postgres{
		zl:   zl,
		pool: pool,
	}, nil
}

func (p *Postgres) Close() error {
	p.pool.Close()
	return nil
}
