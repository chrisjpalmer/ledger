package postgres

import (
	"context"
	"fmt"

	"github.com/chrisjpalmer/ledger/backend/internal/model"
)

func (p *Postgres) GetExpenses(ctx context.Context, month int) ([]model.Expense, error) {
	rows, err := p.pool.Query(ctx, `SELECT
      amount,
			"date",
      id,
			"month",
			"name",
			"paid"
    FROM expenses
    WHERE month = $1;
  `, month)

	if err != nil {
		return nil, fmt.Errorf("error when querying expense %w", err)
	}
	defer rows.Close()

	var ee []model.Expense
	for rows.Next() {
		var exp model.Expense
		if err := rows.Scan(&exp.Amount, &exp.Date, &exp.ID, &exp.Month, &exp.Name, &exp.Paid); err != nil {
			return nil, fmt.Errorf("error while scanning expense %w", err)
		}

		ee = append(ee, exp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning expense %w", err)
	}

	return ee, nil
}

func (p *Postgres) AddExpense(ctx context.Context, expense model.Expense) (string, error) {
	row := p.pool.QueryRow(ctx, `
		INSERT INTO expenses (
			amount,
			"date",
			"month",
			"name",
			"paid"
		) VALUES (
		 	$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id
	`,
		expense.Amount,
		expense.Date,
		expense.Month,
		expense.Name,
		expense.Paid,
	)

	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}
