package ui

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/aviv-baruch/personal-finance-manager/pkg/finance"
)

// HandleCommand processes a single command input
func HandleCommand(input string, fm finance.FinanceManager) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return errors.New("no input provided")
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "add":
		return handleAdd(args, fm)
	case "edit":
		return handleEdit(args, fm)
	default:
		return errors.New("unknown command")
	}
}

// handleAdd parses the arguments for the 'add' command and calls the finance manager
func handleAdd(args []string, fm finance.FinanceManager) error {
	if len(args) != 2 {
		return errors.New("usage: add <amount> <description>")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return errors.New("invalid amount")
	}

	description := args[1]
	return fm.AddTransaction(
		finance.Transaction{
			ID:              fm.ShowOverallItems(),
			Amount:          amount,
			Description:     description,
			TransactionType: 0,
			Date:            time.Now(),
		})
}

func handleEdit(args []string, fm finance.FinanceManager) error {
	if len(args) != 3 {
		return errors.New("usage: add <id> <amount> <description>")
	}

	transactionID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return errors.New("invalid ID")
	}

	updatedAmount, err := strconv.ParseFloat(args[1], 64)
	updatedTransaction := finance.Transaction{
		ID:              transactionID,
		Amount:          updatedAmount,
		Description:     args[2],
		TransactionType: 0,
		Date:            time.Now(),
	}

	fm.EditTransaction(transactionID, updatedTransaction)
	if err != nil {
		return errors.New("Edit went wrong")
	}
	return nil
}

func handleDelete(args []string, fm finance.FinanceManager) error {

	return nil
}

func handleCalculate(args []string, fm finance.FinanceManager) error {

	return nil
}
