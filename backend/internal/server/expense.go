package server

import (
	"context"
	"net/http"

	openapi "github.com/chrisjpalmer/ledger/backend/internal/api/go"
)

func (s *Server) AddExpense(ctx context.Context, month string, expense openapi.Expense) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
func (s *Server) UpdateExpense(ctx context.Context, month string, expenseID string, expense openapi.Expense) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
func (s *Server) DeleteExpense(ctx context.Context, month string, expenseID string) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
