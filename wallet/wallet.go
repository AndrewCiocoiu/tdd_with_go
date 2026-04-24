package wallet

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(deposit_val Bitcoin) {
	w.balance += deposit_val
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(withdraw_amount Bitcoin) error {
	if withdraw_amount > w.balance {
		return errors.New("Not enough money in account")
	}
	w.balance -= withdraw_amount
	return nil
}
