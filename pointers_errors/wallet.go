package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	amount Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	(*w).amount = (*w).amount + deposit
}

func (w *Wallet) Withdraw(amt Bitcoin) {
	w.amount = w.amount - amt
}

func (w *Wallet) Balance() Bitcoin {
	return w.amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
