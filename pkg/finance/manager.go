package main

import "fmt"

type FinanceManager interface {
	AddTransaction(t Transaction) error                // Adds new transaction
	EditTransaction(id int, updated Transaction) error //Edits existing transaction based on ID
	DeleteTransaction(id int) error                    //Delets transaction based on ID
	CalculateBalance() (float64, error)                //calculates the balance from all transactions made
}

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

func (fm *FinanceManagerImpl) CalculateBalance() (float64, error) {
	var sum float64
	if len(fm.Transactions) == 0 {
		return 0, fmt.Errorf("Transaction list is empty")
	}
	for i := 0; i < len(fm.Transactions); i++ {
		sum += fm.Transactions[i].amount
	}

	return sum, nil
}

// This function handles basic error checking for transactions
func ValidateTransaction(t Transaction) error {
	if t.amount <= 0 {
		return fmt.Errorf("Transaction amount cannot be equal or lower than 0")
	}
	if t.description == "" {
		return fmt.Errorf("Transaction description cannot be empty")
	}
	return nil
}
