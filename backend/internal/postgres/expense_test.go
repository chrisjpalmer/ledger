package postgres_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/chrisjpalmer/ledger/backend/internal/model"
	"github.com/chrisjpalmer/ledger/backend/internal/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestPostgres_GetExpenses(t *testing.T) {
	// get config
	cfg := mustTestConfig(t)

	// create db pool
	pgrs, err := postgres.NewTest(zap.NewNop(), cfg)
	require.NoError(t, err, "database connection failed")
	defer pgrs.Close()

	pool := pgrs.Pool()

	type args struct {
		expensesToAdd []model.Expense
	}
	type expenseMonth struct {
		month    int
		expenses []model.Expense
	}
	tests := map[string]struct {
		args    args
		want    []expenseMonth
		wantErr bool
	}{
		"basic": {
			args: args{
				expensesToAdd: []model.Expense{
					{
						Amount: 10.0,
						Date:   date(t, "2025-01-01"),
						Month:  0,
						Name:   "bill",
						Paid:   true,
					},
				},
			},
			want: []expenseMonth{
				{
					expenses: []model.Expense{
						{
							Amount: 10.0,
							Date:   date(t, "2025-01-01"),
							Month:  0,
							Name:   "bill",
							Paid:   true,
						},
					},
					month: 0,
				},
				{month: 1},
			},
		},
		"multiple expense": {
			args: args{
				expensesToAdd: []model.Expense{
					{
						Amount: 10.0,
						Date:   date(t, "2025-01-01"),
						Month:  0,
						Name:   "bill",
						Paid:   true,
					},
					{
						Amount: 5.0,
						Date:   date(t, "2025-01-02"),
						Month:  0,
						Name:   "special bill",
						Paid:   false,
					},
				},
			},
			want: []expenseMonth{
				{
					expenses: []model.Expense{

						{
							Amount: 10.0,
							Date:   date(t, "2025-01-01"),
							Month:  0,
							Name:   "bill",
							Paid:   true,
						},
						{
							Amount: 5.0,
							Date:   date(t, "2025-01-02"),
							Month:  0,
							Name:   "special bill",
							Paid:   false,
						},
					},
					month: 0,
				},
				{month: 1},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// clear database for clean state
			_, err := pool.Exec(context.Background(), "TRUNCATE expenses")
			require.NoError(t, err, "clean database")

			// test query
			for _, i := range tt.args.expensesToAdd {
				_, err := pgrs.AddExpense(context.Background(), i)
				require.NoError(t, err)
			}

			// test expectations
			for _, month := range tt.want {
				got, err := pgrs.GetExpenses(context.Background(), month.month)

				if tt.wantErr {
					require.NotNil(t, err)
				} else {
					require.NoError(t, err, "GetExpenses()")
					requireEqualExpenses(t, month.expenses, got, true)
				}
			}
		})
	}
}
func TestPostgres_AddExpense(t *testing.T) {
	// get config
	cfg := mustTestConfig(t)

	// create db pool
	pgrs, err := postgres.NewTest(zap.NewNop(), cfg)
	require.NoError(t, err, "database connection failed")
	defer pgrs.Close()

	pool := pgrs.Pool()

	type args struct {
		expenses []model.Expense
	}
	tests := map[string]struct {
		args    args
		want    []model.Expense
		wantErr bool
	}{
		"basic": {
			args: args{
				expenses: []model.Expense{
					{
						Amount: 10.0,
						Date:   date(t, "2025-01-01"),
						Month:  0,
						Name:   "bill",
						Paid:   true,
					},
				},
			},
			want: []model.Expense{
				{
					Amount: 10.0,
					Date:   date(t, "2025-01-01"),
					Month:  0,
					Name:   "bill",
					Paid:   true,
				},
			},
		},
		"multiple expense": {
			args: args{
				expenses: []model.Expense{
					{
						Amount: 10.0,
						Date:   date(t, "2025-01-01"),
						Month:  0,
						Name:   "bill",
						Paid:   true,
					},
					{
						Amount: 5.0,
						Date:   date(t, "2025-01-02"),
						Month:  0,
						Name:   "special bill",
						Paid:   false,
					},
				},
			},
			want: []model.Expense{
				{
					Amount: 10.0,
					Date:   date(t, "2025-01-01"),
					Month:  0,
					Name:   "bill",
					Paid:   true,
				},
				{
					Amount: 5.0,
					Date:   date(t, "2025-01-02"),
					Month:  0,
					Name:   "special bill",
					Paid:   false,
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// clear database for clean state
			_, err := pool.Exec(context.Background(), "TRUNCATE expenses")
			require.NoError(t, err, "clean database")

			// test query
			for _, i := range tt.args.expenses {
				gotID, err := pgrs.AddExpense(context.Background(), i)
				if tt.wantErr {
					require.NotNil(t, err)
				} else {
					require.NoError(t, err, "AddExpense()")
				}
				require.NotEmpty(t, gotID)
			}

			// assert conditions
			got := mustGetExpenses(t, pool)

			requireEqualExpenses(t, tt.want, got, true)
		})
	}
}

func mustGetExpenses(t *testing.T, pool *pgxpool.Pool) []model.Expense {
	t.Helper()

	rows, err := pool.Query(context.Background(), `
			SELECT amount,
				date,
        id,
				month,
				name,
				paid
			FROM expenses ORDER BY date`)

	require.NoError(t, err, "get expense")

	defer rows.Close()

	var ii []model.Expense
	for rows.Next() {
		var i model.Expense
		err = rows.Scan(&i.Amount, &i.Date, &i.ID, &i.Month, &i.Name, &i.Paid)
		require.NoError(t, err, "row scan")

		ii = append(ii, i)
	}

	require.NoError(t, rows.Err(), "get expense - rows error")

	return ii
}

func requireEqualExpenses(t *testing.T, want, got []model.Expense, idNotEmpty bool) {
	t.Helper()

	require.Len(t, got, len(want))

	for i, w := range want {
		g := got[i]

		requireEqualExpense(t, w, g, idNotEmpty, strconv.Itoa(i))
	}
}

func requireEqualExpense(t *testing.T, w, g model.Expense, idNotEmpty bool, suffix string) {
	t.Helper()

	require.Equal(t, w.Amount, g.Amount, fmt.Sprintf("%s Amount", suffix))

	require.Equal(t, w.Date, g.Date, fmt.Sprintf("%s Date", suffix))

	if idNotEmpty {
		require.NotEmpty(t, g.ID, fmt.Sprintf("%s ID", suffix))
	} else {
		require.Empty(t, g.ID, fmt.Sprintf("%s ID", suffix))
	}

	require.Equal(t, w.Month, g.Month, fmt.Sprintf("%s Month", suffix))

	require.Equal(t, w.Paid, g.Paid, fmt.Sprintf("%s Paid", suffix))
}
