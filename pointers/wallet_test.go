package goPointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error got nothing")
		}

		if got.Error() != want {
			t.Errorf("Got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw")
	})
}
