package main

import (
	"time"
)

type TransactionType int // type used as enum for transaction type

const (
	Income  TransactionType = iota // iota starts at 0
	Expense                        // implicitly Income + 1
)

type Transaction struct {
	id              int
	amount          float64
	description     string
	transactionType TransactionType
	date            time.Time
}
