package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aviv-baruch/personal-finance-manager/pkg/finance"
	"github.com/aviv-baruch/personal-finance-manager/pkg/utils"
)

// HandleCommand processes a single command input
func HandleCommand(input string, fm finance.FinanceManager) error {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return &utils.ValidationError{
			Message: "No input received",
			ErrCode: utils.InvalidUsage,
		}
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "add":
		return handleAdd(args, fm)
	case "edit":
		return handleEdit(args, fm)
	case "delete":
		return handleDelete(args, fm)
	case "calculate":
		return handleCalculate(args, fm)
	default:
		return &utils.ValidationError{
			Message: "Unkown Command",
			ErrCode: utils.InvalidUsage,
		}
	}
}

// handleAdd parses the arguments for the 'add' command and calls the finance manager
func handleAdd(args []string, fm finance.FinanceManager) error {
	if len(args) != 2 {
		return &utils.ValidationError{
			Message: "usage: add <amount> <description>",
			ErrCode: utils.InvalidUsage,
		}
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return &utils.ValidationError{
			Message: "Invalid Amount entered",
			ErrCode: utils.InvalidData,
		}
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
		return &utils.ValidationError{
			Message: "usage: add <id> <amount> <description>",
			ErrCode: utils.InvalidData,
		}
	}

	transactionID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return &utils.ValidationError{
			Message: "Invalid ID",
			ErrCode: utils.InvalidData,
		}
	}

	fmt.Printf("Editing transacton ID %d with amount %f with description %s\n", transactionID, fm.GetTransactions()[transactionID].Amount, fm.GetTransactions()[transactionID].Description)
	updatedAmount, _ := strconv.ParseFloat(args[1], 64)
	updatedTransaction := finance.Transaction{
		ID:              transactionID,
		Amount:          updatedAmount,
		Description:     args[2],
		TransactionType: 0,
		Date:            time.Now(),
	}

	return fm.EditTransaction(transactionID, updatedTransaction)
}

func handleDelete(args []string, fm finance.FinanceManager) error {
	if len(args) != 1 {
		return &utils.ValidationError{
			Message: "usage: delete <id>",
			ErrCode: utils.InvalidData,
		}
	}

	transactionID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return &utils.ValidationError{
			Message: "Invalid ID",
			ErrCode: utils.InvalidData,
		}
	}

	return fm.DeleteTransaction(transactionID)
}

func handleCalculate(args []string, fm finance.FinanceManager) error {
	if len(args) > 0 {
		return &utils.ValidationError{
			Message: "usage: calculate",
			ErrCode: utils.InvalidData,
		}
	}
	balance, err := fm.CalculateBalance()
	if err != nil {
		return &utils.ValidationError{
			Message: "Transaction list is empty",
			ErrCode: utils.InvalidData,
		}

	} else {
		fmt.Printf("Current balance is: $%.2f\n", balance) // Print the balance
		return nil
	}
}
