package wallet

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(deposit_val int) {
	println("Deposit complete")
}

func (w Wallet) Balance() int {
	return 0
}
