package postgres

import (
	"context"
	"fmt"

	"github.com/chrisjpalmer/ledger/backend/internal/model"
)

func (p *Postgres) GetIncome(ctx context.Context, month int) ([]model.Income, error) {
	rows, err := p.pool.Query(ctx, `SELECT
      amount,
			"date",
      id,
			"month",
			"name",
			"received"
    FROM income
    WHERE month = $1;
  `, month)

	if err != nil {
		return nil, fmt.Errorf("error when querying income %w", err)
	}
	defer rows.Close()

	var ii []model.Income
	for rows.Next() {
		var inc model.Income
		if err := rows.Scan(&inc.Amount, &inc.Date, &inc.ID, &inc.Month, &inc.Name, &inc.Received); err != nil {
			return nil, fmt.Errorf("error while scanning income %w", err)
		}

		ii = append(ii, inc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning income %w", err)
	}

	return ii, nil
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
