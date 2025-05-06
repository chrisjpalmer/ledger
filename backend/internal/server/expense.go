package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	openapi "github.com/chrisjpalmer/ledger/backend/internal/api/go"
	"github.com/chrisjpalmer/ledger/backend/internal/model"
	"go.uber.org/zap"
)

func (s *Server) GetExpenses(ctx context.Context, month int32) (openapi.ImplResponse, error) {
	exp, err := s.pgs.GetExpenses(ctx, int(month))
	if err != nil {
		s.zl.Error("internal error while getting expenses", zap.Error(err))
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{
			Msg: "error when fetching expenses",
		}), nil
	}

	return openapi.Response(http.StatusOK, openapi.GetExpensesResponse{
		Expenses: mapExpenses(exp),
	}), nil
}

func mapExpenses(inc []model.Expense) []openapi.Expense {
	rsInc := make([]openapi.Expense, 0, len(inc))
	for _, in := range inc {
		rsInc = append(rsInc, openapi.Expense{
			Amount: in.Amount,
			Date:   in.Date.Format(time.RFC3339),
			Id:     in.ID,
			Name:   in.Name,
			Paid:   in.Paid,
		})
	}

	return rsInc
}

func (s *Server) AddExpense(ctx context.Context, month int32, expense openapi.Expense) (openapi.ImplResponse, error) {
	date, err := parseDate(expense.Date)
	if err != nil {
		return openapi.Response(http.StatusBadRequest, openapi.ErrorResponse{
			Msg: fmt.Sprintf("error when parsing url parameter `date`: %s", err.Error()),
		}), nil
	}

	exp := model.Expense{
		Amount: expense.Amount,
		Date:   date,
		Month:  int(month),
		Name:   expense.Name,
		Paid:   expense.Paid,
	}

	id, err := s.pgs.AddExpense(ctx, exp)
	if err != nil {
		s.zl.Error("internal error while adding an expense", zap.Error(err))
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{
			Msg: "error when adding an expense",
		}), nil
	}

	return openapi.Response(http.StatusOK, openapi.ExpenseResponse{
		Id: id,
	}), nil

}
func (s *Server) UpdateExpense(ctx context.Context, month int32, expenseID string, expense openapi.Expense) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
func (s *Server) DeleteExpense(ctx context.Context, month int32, expenseID string) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
