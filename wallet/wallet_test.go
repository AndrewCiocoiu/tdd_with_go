package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Balance is deposited correctly", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := 10

		if want != got {
			t.Errorf("Wanted %d, got %d", want, got)
		}

	})
}
