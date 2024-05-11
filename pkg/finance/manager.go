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
	if t.amount <= 0 {
		return fmt.Errorf("Transaction amount cannot be equal or lower than 0")
	}
	if t.description == "" {
		return fmt.Errorf("Transaction description cannot be empty")
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
		return fmt.Errorf("Transaction id should be positive number")
	}

	for i := 0; i < len(fm.Transactions); i++ {
		if fm.Transactions[i].id == id {
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
		if fm.Transactions[i].id == id {
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
		if fm.Transactions[i].transactionType == Income {
			balance += fm.Transactions[i].amount
		}
		if fm.Transactions[i].transactionType == Expense {
			balance -= fm.Transactions[i].amount
		}
	}
	return balance, nil
}

// This function handles basic error checking for transactions
// TODO: move to errors
func ValidateTransaction(t Transaction) error {
	if t.transactionType > 1 || t.transactionType < 0 {
		return fmt.Errorf("Couldn't find transaction type")
	}
	if t.id <= 0 {
		return fmt.Errorf("ID cannot be negative")
	}
	if t.amount <= 0 {
		return fmt.Errorf("Transaction amount cannot be equal or lower than 0")
	}
	if t.description == "" {
		return fmt.Errorf("Transaction description cannot be empty")
	}
	return nil
}
