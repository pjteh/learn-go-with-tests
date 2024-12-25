package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	amount Bitcoin
}

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(deposit Bitcoin) {
	(*w).amount = (*w).amount + deposit
}

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.amount {
		return ErrInsufficientFunds
	}
	w.amount = w.amount - amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
