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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// LedgerAPIController binds http requests to an api service and writes the service results to the http response
type LedgerAPIController struct {
	service LedgerAPIServicer
	errorHandler ErrorHandler
}

// LedgerAPIOption for how the controller is set up.
type LedgerAPIOption func(*LedgerAPIController)

// WithLedgerAPIErrorHandler inject ErrorHandler into controller
func WithLedgerAPIErrorHandler(h ErrorHandler) LedgerAPIOption {
	return func(c *LedgerAPIController) {
		c.errorHandler = h
	}
}

// NewLedgerAPIController creates a default api controller
func NewLedgerAPIController(s LedgerAPIServicer, opts ...LedgerAPIOption) *LedgerAPIController {
	controller := &LedgerAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the LedgerAPIController
func (c *LedgerAPIController) Routes() Routes {
	return Routes{
		"AddIncome": Route{
			strings.ToUpper("Post"),
			"/month/{month}/income",
			c.AddIncome,
		},
		"UpdateIncome": Route{
			strings.ToUpper("Put"),
			"/month/{month}/income/{incomeId}",
			c.UpdateIncome,
		},
		"DeleteIncome": Route{
			strings.ToUpper("Delete"),
			"/month/{month}/income/{incomeId}",
			c.DeleteIncome,
		},
		"AddExpense": Route{
			strings.ToUpper("Post"),
			"/month/{month}/expense",
			c.AddExpense,
		},
		"UpdateExpense": Route{
			strings.ToUpper("Put"),
			"/month/{month}/expense_id",
			c.UpdateExpense,
		},
		"DeleteExpense": Route{
			strings.ToUpper("Delete"),
			"/month/{month}/expense_id",
			c.DeleteExpense,
		},
	}
}

// AddIncome - Add a new line of income
func (c *LedgerAPIController) AddIncome(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monthParam := params["month"]
	if monthParam == "" {
		c.errorHandler(w, r, &RequiredError{"month"}, nil)
		return
	}
	var incomeParam Income
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&incomeParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIncomeRequired(incomeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertIncomeConstraints(incomeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddIncome(r.Context(), monthParam, incomeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateIncome - Update a line of income
func (c *LedgerAPIController) UpdateIncome(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monthParam := params["month"]
	if monthParam == "" {
		c.errorHandler(w, r, &RequiredError{"month"}, nil)
		return
	}
	incomeIdParam := params["incomeId"]
	if incomeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"incomeId"}, nil)
		return
	}
	var incomeParam Income
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&incomeParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertIncomeRequired(incomeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertIncomeConstraints(incomeParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateIncome(r.Context(), monthParam, incomeIdParam, incomeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteIncome - Delete a line of income
func (c *LedgerAPIController) DeleteIncome(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monthParam := params["month"]
	if monthParam == "" {
		c.errorHandler(w, r, &RequiredError{"month"}, nil)
		return
	}
	incomeIdParam := params["incomeId"]
	if incomeIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"incomeId"}, nil)
		return
	}
	result, err := c.service.DeleteIncome(r.Context(), monthParam, incomeIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// AddExpense - Add a new expense line
func (c *LedgerAPIController) AddExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monthParam := params["month"]
	if monthParam == "" {
		c.errorHandler(w, r, &RequiredError{"month"}, nil)
		return
	}
	var expenseParam Expense
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&expenseParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertExpenseRequired(expenseParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertExpenseConstraints(expenseParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddExpense(r.Context(), monthParam, expenseParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateExpense - Update an expense line
func (c *LedgerAPIController) UpdateExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monthParam := params["month"]
	if monthParam == "" {
		c.errorHandler(w, r, &RequiredError{"month"}, nil)
		return
	}
	expenseIdParam := params["expenseId"]
	if expenseIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"expenseId"}, nil)
		return
	}
	var expenseParam Expense
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&expenseParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertExpenseRequired(expenseParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertExpenseConstraints(expenseParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateExpense(r.Context(), monthParam, expenseIdParam, expenseParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteExpense - Delete an expense line
func (c *LedgerAPIController) DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monthParam := params["month"]
	if monthParam == "" {
		c.errorHandler(w, r, &RequiredError{"month"}, nil)
		return
	}
	expenseIdParam := params["expenseId"]
	if expenseIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"expenseId"}, nil)
		return
	}
	result, err := c.service.DeleteExpense(r.Context(), monthParam, expenseIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
