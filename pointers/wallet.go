package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("insufficient funds")
	}
	w.balance -= amount
	return nil
}
