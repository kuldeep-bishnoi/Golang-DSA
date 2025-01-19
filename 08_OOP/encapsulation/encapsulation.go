package encapsulation

import "fmt"

// BankAccount struct demonstrating encapsulation
type BankAccount struct {
	balance       float64 // Public field
	accountNumber string  // Private field
}

// NewBankAccount creates a new BankAccount
func NewBankAccount(accountNumber string, initialBalance float64) *BankAccount {
	return &BankAccount{balance: initialBalance, accountNumber: accountNumber}
}

// Deposit adds money to the account
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.balance += amount
	}
}

// Withdraw removes money from the account
func (b *BankAccount) Withdraw(amount float64) {
	if amount > 0 && amount <= b.balance {
		b.balance -= amount
	}
}

// GetBalance returns the current balance
func (b *BankAccount) GetBalance() float64 {
	return b.balance
}

func main() {
	account := NewBankAccount("123456789", 1000.0)
	account.Deposit(500.0)
	account.Withdraw(200.0)
	fmt.Println("Current Balance:", account.GetBalance())
}
