package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/chrisjpalmer/ledger/backend/internal/model"
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
	return newPostgres(zl, cfg,
		10,
		100,
		time.Minute*5,
		time.Minute*1,
	)
}

// NewTest - creates a postgres with a pool size suitable for testing
func NewTest(zl *zap.Logger, cfg Config) (*Postgres, error) {
	return newPostgres(zl, cfg,
		1,
		5,
		time.Second*5,
		time.Second*1,
	)
}

func newPostgres(zl *zap.Logger, cfg Config, minConns, maxConns int32, maxConnLifetime, maxConnIdleTime time.Duration) (*Postgres, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	c, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("error while parsing config: %w", err)
	}

	c.MinConns = minConns
	c.MaxConns = maxConns
	c.MaxConnLifetime = maxConnLifetime
	c.MaxConnIdleTime = maxConnIdleTime

	pool, err := pgxpool.NewWithConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("error while creating pgx pool %w", err)
	}

	return &Postgres{
		zl:   zl,
		pool: pool,
	}, nil
}

func (p *Postgres) AddIncome(ctx context.Context, income model.Income) (string, error) {
	row := p.pool.QueryRow(ctx, `
		INSERT INTO income (
			amount,
			"date",
			"month",
			"name",
			"received"
		) VALUES (
		 	$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id
	`,
		income.Amount,
		income.Date,
		income.Month,
		income.Name,
		income.Received,
	)

	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

// Pool - returns the underlying pgxpool
func (p *Postgres) Pool() *pgxpool.Pool {
	return p.pool
}

func (p *Postgres) Close() error {
	p.pool.Close()
	return nil
}
