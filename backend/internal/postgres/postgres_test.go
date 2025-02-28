package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/chrisjpalmer/ledger/backend/config"
	"github.com/chrisjpalmer/ledger/backend/internal/model"
	"github.com/chrisjpalmer/ledger/backend/internal/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestPostgres_AddIncome(t *testing.T) {
	// get config
	cfg := mustTestConfig(t)

	// create db pool
	pgrs, err := postgres.NewTest(zap.NewNop(), cfg)
	require.NoError(t, err, "database connection failed")
	defer pgrs.Close()

	pool := pgrs.Pool()

	type args struct {
		income []model.Income
	}
	tests := map[string]struct {
		args    args
		want    []model.Income
		wantErr bool
	}{
		"basic": {
			args: args{
				income: []model.Income{
					{
						Amount:   10.0,
						Date:     date(t, "2025-01-01"),
						Month:    0,
						Name:     "salary",
						Received: true,
					},
				},
			},
			want: []model.Income{
				{
					Amount:   10.0,
					Date:     date(t, "2025-01-01"),
					Month:    0,
					Name:     "salary",
					Received: true,
				},
			},
		},
		"multiple income": {
			args: args{
				income: []model.Income{
					{
						Amount:   10.0,
						Date:     date(t, "2025-01-01"),
						Month:    0,
						Name:     "salary",
						Received: true,
					},
					{
						Amount:   5.0,
						Date:     date(t, "2025-01-02"),
						Month:    0,
						Name:     "bonus",
						Received: false,
					},
				},
			},
			want: []model.Income{
				{
					Amount:   10.0,
					Date:     date(t, "2025-01-01"),
					Month:    0,
					Name:     "salary",
					Received: true,
				},
				{
					Amount:   5.0,
					Date:     date(t, "2025-01-02"),
					Month:    0,
					Name:     "bonus",
					Received: false,
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// clear database for clean state
			_, err := pool.Exec(context.Background(), "TRUNCATE income")
			require.NoError(t, err, "clean database")

			// test query
			for _, i := range tt.args.income {
				gotID, err := pgrs.AddIncome(context.Background(), i)
				if tt.wantErr {
					require.NotNil(t, err)
				} else {
					require.NoError(t, err, "AddIncome()")
				}
				require.NotEmpty(t, gotID)
			}

			// assert conditions
			got := mustGetIncome(t, pool)

			require.Equal(t, tt.want, got)
		})
	}
}

func date(t *testing.T, s string) time.Time {
	t.Helper()

	tm, err := time.Parse(time.DateOnly, s)
	require.NoError(t, err, "parsing time")

	return tm
}

func mustGetIncome(t *testing.T, pool *pgxpool.Pool) []model.Income {
	t.Helper()

	rows, err := pool.Query(context.Background(), `
			SELECT amount,
				date,
				month,
				name,
				received
			FROM income ORDER BY date`)

	require.NoError(t, err, "get income")

	defer rows.Close()

	var ii []model.Income
	for rows.Next() {
		var i model.Income
		err = rows.Scan(&i.Amount, &i.Date, &i.Month, &i.Name, &i.Received)
		require.NoError(t, err, "row scan")

		ii = append(ii, i)
	}

	require.NoError(t, rows.Err(), "get income - rows error")

	return ii
}

func mustTestConfig(t *testing.T) postgres.Config {
	t.Helper()

	if config.HasDotEnv("../../") {
		config.LoadDotEnv("../../")
	}

	var e config.Errors
	cfg := config.LoadPostgresConfig(&e)

	var errs []string
	e.ForEach(func(err error) { errs = append(errs, err.Error()) })

	require.False(t, e.HasErrors(), errs)

	return cfg
}
