package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit test", func(t *testing.T) {
		wallet := Wallet{}
		depositQuantity := Bitcoin(10)
		balance := Bitcoin(10)

		wallet.Deposit(depositQuantity)
		assertBalance(t, wallet, balance)

	})

	t.Run("withdrawal test, more than deposited", func(t *testing.T) {
		wallet := Wallet{}
		depositQuantity := Bitcoin(10)
		balance := Bitcoin(10)

		wallet.Deposit(depositQuantity)
		err := wallet.Withdraw(20)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, balance)
	})

	t.Run("withdrawal test, less than deposited", func(t *testing.T) {
		wallet := Wallet{}
		depositQuantity := Bitcoin(30)
		finalBalance := 10
		balance := Bitcoin(finalBalance)

		wallet.Deposit(depositQuantity)
		err := wallet.Withdraw(20)

		assertError(t, err, ErrNoError)
		assertBalance(t, wallet, balance)
	})
}

// Compares the given balance against the desired balance
func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()

	got := w.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// Checks for an error
func assertError(t testing.TB, err error, desiredError error) {
	t.Helper()
	want := desiredError.Error()

	if err == nil {
		t.Fatal("wanted an error but none was given")
	}

	if err.Error() != want {
		t.Errorf("got %q, want %q", err, desiredError)
	}
}
