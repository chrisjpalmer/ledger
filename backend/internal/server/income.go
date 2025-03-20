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

func (s *Server) GetIncome(ctx context.Context, month int32) (openapi.ImplResponse, error) {
	inc, err := s.pgs.GetIncome(ctx, int(month))
	if err != nil {
		s.zl.Error("internal error while getting income", zap.Error(err))
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{
			Msg: "error when fetching income",
		}), nil
	}

	return openapi.Response(http.StatusOK, openapi.GetIncomeResponse{
		Income: mapIncome(inc),
	}), nil
}

func mapIncome(inc []model.Income) []openapi.Income {
	rsInc := make([]openapi.Income, 0, len(inc))
	for _, in := range inc {
		rsInc = append(rsInc, openapi.Income{
			Amount:   in.Amount,
			Date:     in.Date.Format(time.RFC3339),
			Id:       in.ID,
			Name:     in.Name,
			Received: in.Received,
		})
	}

	return rsInc
}

func (s *Server) AddIncome(ctx context.Context, month int32, income openapi.Income) (openapi.ImplResponse, error) {
	date, err := parseDate(income.Date)
	if err != nil {
		return openapi.Response(http.StatusBadRequest, openapi.ErrorResponse{
			Msg: fmt.Sprintf("error when parsing url parameter `date`: %s", err.Error()),
		}), nil
	}

	inc := model.Income{
		Amount:   income.Amount,
		Date:     date,
		Month:    int(month),
		Name:     income.Name,
		Received: income.Received,
	}

	id, err := s.pgs.AddIncome(ctx, inc)
	if err != nil {
		s.zl.Error("internal error while adding income", zap.Error(err))
		return openapi.Response(http.StatusInternalServerError, openapi.ErrorResponse{
			Msg: "error when adding income",
		}), nil
	}

	return openapi.Response(http.StatusOK, openapi.IncomeResponse{
		Id: id,
	}), nil
}

func (s *Server) UpdateIncome(ctx context.Context, month int32, incomeID string, income openapi.Income) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func (s *Server) DeleteIncome(ctx context.Context, month int32, incomeID string) (openapi.ImplResponse, error) {
	return openapi.Response(http.StatusNotImplemented, nil), nil
}

func parseDate(d string) (time.Time, error) {
	return time.Parse(time.DateOnly, d)
}
