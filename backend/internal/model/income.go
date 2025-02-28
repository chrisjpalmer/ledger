package model

import "time"

// Income - represents an income entry
type Income struct {
	Amount   float32
	Date     time.Time
	ID       string
	Month    int
	Name     string
	Received bool
}
