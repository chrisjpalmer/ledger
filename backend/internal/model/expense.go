package model

import "time"

// Expense - represents an expense entry
type Expense struct {
	Amount float32
	Date   time.Time
	ID     string
	Month  int
	Name   string
	Paid   bool
}
