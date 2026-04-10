package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Balance is deposited correctly", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(11)

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
}
