package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	}
	t.Run("deposit method", func(t *testing.T) {
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw method", func(t *testing.T) {
		wallet.Withdraw(10)
		want := Bitcoin(0)
		assertBalance(t, wallet, want)
	})
}
