package finance

import (
	"fmt"

	"github.com/aviv-baruch/personal-finance-manager/pkg/utils"
)

// App's core
type FinanceManager interface {
	GetTransactions() []Transaction                      // Returns transactions arr
	ShowOverallItems() int64                             // Returns amount of items in Transactions array
	AddTransaction(t Transaction) error                  // Adds new transaction
	EditTransaction(id int64, updated Transaction) error //Edits existing transaction based on ID
	DeleteTransaction(id int64) error                    //Delets transaction based on ID
	CalculateBalance() (float64, error)                  //calculates the balance from all transactions made
}

// App's core
type FinanceManagerImpl struct {
	Transactions []Transaction
	OverallItems int64
}

func NewFinanceManagerImpl() *FinanceManagerImpl {
	return &FinanceManagerImpl{
		Transactions: []Transaction{},
		OverallItems: 0,
	}
}

func (fm FinanceManagerImpl) ShowOverallItems() int64 {
	return fm.OverallItems
}

func (fm FinanceManagerImpl) GetTransactions() []Transaction {
	return fm.Transactions
}

// Adds a new transaction
func (fm *FinanceManagerImpl) AddTransaction(t Transaction) error {
	if t.Amount <= 0 {
		return &utils.ValidationError{
			Message: "Transaction amount must be positive",
			ErrCode: utils.GeneralError,
		}
	}
	if t.Description == "" {
		return &utils.ValidationError{
			Message: "Transaction description cannot be empty",
			ErrCode: utils.GeneralError,
		}
	}
	fm.Transactions = append(fm.Transactions, t)
	fm.OverallItems++

	fmt.Printf("Added $%.2f due to %s on %s\n", t.Amount, t.Description, t.Date.Format("2006-01-02"))
	fmt.Printf("Transaction id: %d\n", t.ID)

	balance, _ := fm.CalculateBalance()                // Call CalculateBalance method on fm object
	fmt.Printf("Current balance is: $%.2f\n", balance) // Print the balance
	return nil
}

// Edits existing transaction
func (fm *FinanceManagerImpl) EditTransaction(id int64, updated Transaction) error {
	if err := ValidateTransaction(updated); err != nil {
		return err
	}

	for i := 0; i < len(fm.Transactions); i++ { //TODO
		if fm.Transactions[i].ID == id {
			fm.Transactions[i] = updated
			fmt.Printf("new values of transacton ID %d with amount %f with description %s\n", fm.Transactions[i].ID, fm.Transactions[i].Amount, fm.Transactions[i].Description)

			balance, _ := fm.CalculateBalance()
			fmt.Printf("Current balance is: $%.2f\n", balance) // Print the balance
			return nil
		}
	}

	return &utils.ValidationError{
		Message: fmt.Sprintf("Found no transaction with ID %d", id),
		ErrCode: utils.GeneralError,
	}

}

// This function used to delete existing transaction out of the ones belongs to the finance manager
// Recieves an int (which is the ID of the transaction) and return error if exists
func (fm *FinanceManagerImpl) DeleteTransaction(id int64) error {
	if id < 0 {
		return &utils.ValidationError{
			Message: "Transaction id should be positive number",
			ErrCode: utils.InvalidData}
	}

	for i := 0; i < len(fm.Transactions); i++ {
		if fm.Transactions[i].ID == id {
			fm.Transactions = append(fm.Transactions[:i], fm.Transactions[i+1:]...) //slice around the elemnt

			fmt.Printf("Deleted transaciton ID #%d\n", id) // Print the balance
			balance, _ := fm.CalculateBalance()
			fmt.Printf("Current balance is: $%.2f\n", balance) // Print the balance
			return nil
		}
	}

	return &utils.ValidationError{
		Message: fmt.Sprintf("Found no transaction with ID %d", id),
		ErrCode: utils.GeneralError,
	}
}

// This function used to calculate balance
// Returns the balance as float64 and error message if exists.
func (fm *FinanceManagerImpl) CalculateBalance() (float64, error) {
	var balance float64
	if len(fm.Transactions) == 0 {
		return 0, &utils.ValidationError{
			Message: "Transaction list is empty",
			ErrCode: utils.GeneralError,
		}
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
		return &utils.ValidationError{
			Message: "Couldn't find transaction type",
			ErrCode: utils.InvalidData}
	}
	if t.ID < 0 {
		return &utils.ValidationError{
			Message: "ID cannot be negative",
			ErrCode: utils.InvalidData}
	}
	if t.Amount <= 0 {
		return &utils.ValidationError{
			Message: "Transaction amount cannot be equal or lower than 0",
			ErrCode: utils.InvalidData}
	}
	if t.Description == "" {
		return &utils.ValidationError{
			Message: "Transaction description cannot be empty",
			ErrCode: utils.InvalidData}
	}
	return nil
}
