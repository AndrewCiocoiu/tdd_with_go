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
}
