package finance

import (
	"time"
)

type TransactionType int // type used as enum for transaction type

const (
	Income  TransactionType = iota // iota starts at 0
	Expense                        // implicitly Income + 1
)

// Type used to represent each one of the transactions
// Transactions exists as part of array of the balance manager, as can be seen on manager.go
type Transaction struct {
	ID              int64
	Amount          float64
	Description     string
	TransactionType TransactionType
	Date            time.Time
}
