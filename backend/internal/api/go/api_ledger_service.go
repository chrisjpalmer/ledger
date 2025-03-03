// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Ledger API
 *
 * This is the API for the ledger backend.
 *
 * API version: 0.0.1
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// LedgerAPIService is a service that implements the logic for the LedgerAPIServicer
// This service should implement the business logic for every endpoint for the LedgerAPI API.
// Include any external packages or services that will be required by this service.
type LedgerAPIService struct {
}

// NewLedgerAPIService creates a default api service
func NewLedgerAPIService() *LedgerAPIService {
	return &LedgerAPIService{}
}

// AddIncome - Add a new line of income
func (s *LedgerAPIService) AddIncome(ctx context.Context, month int32, income Income) (ImplResponse, error) {
	// TODO - update AddIncome with the required logic for this service method.
	// Add api_ledger_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, IncomeResponse{}) or use other options such as http.Ok ...
	// return Response(200, IncomeResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddIncome method not implemented")
}

// UpdateIncome - Update a line of income
func (s *LedgerAPIService) UpdateIncome(ctx context.Context, month int32, incomeId string, income Income) (ImplResponse, error) {
	// TODO - update UpdateIncome with the required logic for this service method.
	// Add api_ledger_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, IncomeResponse{}) or use other options such as http.Ok ...
	// return Response(200, IncomeResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(404, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateIncome method not implemented")
}

// DeleteIncome - Delete a line of income
func (s *LedgerAPIService) DeleteIncome(ctx context.Context, month int32, incomeId string) (ImplResponse, error) {
	// TODO - update DeleteIncome with the required logic for this service method.
	// Add api_ledger_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, IncomeResponse{}) or use other options such as http.Ok ...
	// return Response(200, IncomeResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(404, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteIncome method not implemented")
}

// AddExpense - Add a new expense line
func (s *LedgerAPIService) AddExpense(ctx context.Context, month int32, expense Expense) (ImplResponse, error) {
	// TODO - update AddExpense with the required logic for this service method.
	// Add api_ledger_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExpenseResponse{}) or use other options such as http.Ok ...
	// return Response(200, ExpenseResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddExpense method not implemented")
}

// UpdateExpense - Update an expense line
func (s *LedgerAPIService) UpdateExpense(ctx context.Context, month int32, expenseId string, expense Expense) (ImplResponse, error) {
	// TODO - update UpdateExpense with the required logic for this service method.
	// Add api_ledger_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExpenseResponse{}) or use other options such as http.Ok ...
	// return Response(200, ExpenseResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(404, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateExpense method not implemented")
}

// DeleteExpense - Delete an expense line
func (s *LedgerAPIService) DeleteExpense(ctx context.Context, month int32, expenseId string) (ImplResponse, error) {
	// TODO - update DeleteExpense with the required logic for this service method.
	// Add api_ledger_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExpenseResponse{}) or use other options such as http.Ok ...
	// return Response(200, ExpenseResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(404, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteExpense method not implemented")
}
