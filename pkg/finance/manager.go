package main

import "fmt"

// App's core
type FinanceManager interface {
	AddTransaction(t Transaction) error                // Adds new transaction
	EditTransaction(id int, updated Transaction) error //Edits existing transaction based on ID
	DeleteTransaction(id int) error                    //Delets transaction based on ID
	CalculateBalance() (float64, error)                //calculates the balance from all transactions made
}

// App's core
type FinanceManagerImpl struct {
	Transactions []Transaction
}

// Adds a new transaction
func (fm *FinanceManagerImpl) AddTransaction(t Transaction) error {
	if t.Amount <= 0 {
		return &ValidationError{Message: "Transaction amount must be positive", ErrCode: GeneralError}
	}
	if t.Description == "" {
		return &ValidationError{Message: "Transaction description cannot be empty", ErrCode: GeneralError}
	}
	fm.Transactions = append(fm.Transactions, t)
	return nil
}

// Edits existing transaction
func (fm *FinanceManagerImpl) EditTransaction(id int, updated Transaction) error {
	if err := ValidateTransaction(updated); err != nil {
		return err
	}
	if id <= 0 {
		return &ValidationError{Message: "Transaction id should be positive number", ErrCode: GeneralError}
	}

	for i := 0; i < len(fm.Transactions); i++ {
		if fm.Transactions[i].ID == id {
			fm.Transactions[i] = updated
			return nil
		}
	}

	return fmt.Errorf("Found no transaction with ID %d", id)
}

// This function used to delete existing transaction out of the ones belongs to the finance manager
// Recieves an int (which is the ID of the transaction) and return error if exists
func (fm *FinanceManagerImpl) DeleteTransaction(id int) error {
	if id <= 0 {
		return fmt.Errorf("Transaction id should be positive number")
	}

	for i := 0; i < len(fm.Transactions); i++ {
		if fm.Transactions[i].ID == id {
			fm.Transactions = append(fm.Transactions[:i], fm.Transactions[i+1:]...) //slice around the elemnt
			return nil
		}
	}

	return fmt.Errorf("Found no transaction with ID %d", id)

}

// This function used to calculate balance
// Returns the balance as float64 and error message if exists.
func (fm *FinanceManagerImpl) CalculateBalance() (float64, error) {
	var balance float64
	if len(fm.Transactions) == 0 {
		return 0, fmt.Errorf("Transaction list is empty")
	}
	for i := 0; i < len(fm.Transactions); i++ {
		if fm.Transactions[i].TransactionType == Income {
			balance += fm.Transactions[i].Amount
		}
		if fm.Transactions[i].TransactionType == Expense {
			balance -= fm.Transactions[i].Amount
		}
	}
	return balance, nil
}

// This function handles basic error checking for transactions
// TODO: move to errors
func ValidateTransaction(t Transaction) error {
	if t.TransactionType > 1 || t.TransactionType < 0 {
		return &ValidationError{Message: "Couldn't find transaction type", ErrCode: GeneralError}
	}
	if t.ID <= 0 {
		return &ValidationError{Message: "ID cannot be negative", ErrCode: GeneralError}
	}
	if t.Amount <= 0 {
		return &ValidationError{Message: "Transaction amount cannot be equal or lower than 0", ErrCode: GeneralError}
	}
	if t.Description == "" {
		return &ValidationError{Message: "Transaction description cannot be empty", ErrCode: GeneralError}
	}
	return nil
}
