package main

import (
	"errors"
	"fmt"
)

// Stringer to string
type Stringer interface {
	String() string
}

// Bitcoin type
type Bitcoin int

// String returns printable string
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet type holds bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit money in the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns funds in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds error for Withdraw method
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw withdraws funds
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
