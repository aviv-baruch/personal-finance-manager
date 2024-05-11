package main

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
	id              int
	amount          float64
	description     string
	transactionType TransactionType
	date            time.Time
}
