package server

import (
	"context"
	"net/http"

	openapi "github.com/chrisjpalmer/ledger/backend/internal/api/go"
)

func (s *Server) AddIncome(ctx context.Context, month int32, income openapi.Income) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func (s *Server) UpdateIncome(ctx context.Context, month int32, incomeID string, income openapi.Income) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func (s *Server) DeleteIncome(ctx context.Context, month int32, incomeID string) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}
