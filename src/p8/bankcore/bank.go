package bank

import (
	"errors"
	"fmt"
)

func Hello() string {
	return "Hey! I'm working!"
}

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// Bank ...
type Bank interface {
	Statement() string
}

// Deposit ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

func (a *Account) Transfer(amount float64, ac *Account) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	if err := ac.Deposit(amount); err != nil {
		return err
	}
	return nil
}

// Statement ...
func Statement(b Bank) string {
	return b.Statement()
}

func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}
