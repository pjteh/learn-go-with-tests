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
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
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
	t.Run("withdraw insufficient fundsmethod", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(100)
		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})
}
