package pointers_and_errors

import "errors"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount int) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

var ErrInsufficientFunds = errors.New("insufficient funds")
