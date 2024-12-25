package wallet

type Wallet struct {
	Amount int
}

func (w *Wallet) Deposit(deposit int) {
	(*w).Amount = (*w).Amount + deposit
}

func (w *Wallet) Balance() int {
	return w.Amount
}
