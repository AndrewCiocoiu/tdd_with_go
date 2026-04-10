package wallet

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(deposit_val int) {
	w.balance += deposit_val
}

func (w *Wallet) Balance() int {
	return w.balance
}
