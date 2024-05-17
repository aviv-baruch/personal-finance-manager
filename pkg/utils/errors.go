package utils

import (
	"fmt"
)

type ErrorCode int // type used as enum for transaction type

const (
	GeneralError ErrorCode = iota // iota starts at 0
	InvalidUsage
	InvalidData
)

type ValidationError struct {
	Message string
	ErrCode ErrorCode
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.ErrCode, e.Message)
}
