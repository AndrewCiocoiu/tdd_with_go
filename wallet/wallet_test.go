package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Balance is deposited correctly", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if want != got {
			t.Errorf("Wanted %s, got %s", want, got)
		}

	})

	t.Run("Withdraw correctly", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("Wanted %s, got %s", want, got)
		}

	})

	t.Run("Witdraw Overdraft", func(t *testing.T) {
		starting_balance := Bitcoin(10)

		wallet := Wallet{balance: starting_balance}

		err := wallet.Withdraw(Bitcoin(100))

		got := wallet.Balance()
		want := starting_balance

		if got != want {
			t.Errorf("Expected starting balance to stay: %s, got %s", want, got)
		}

		if err == nil {
			t.Errorf("Wanted an error but didn't get one")
		}

	})
}
