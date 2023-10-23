package pointerserrors

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{}

	w.Deposit(Bitcoin(10))

	got := w.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestWalletStringer(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(20)}
		err := w.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		startingbalance := Bitcoin(10)
		w := Wallet{balance: startingbalance}
		err := w.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, startingbalance)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("Wanted %q got %q", want, got)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Didn't want an error but got one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("Want %q, got %q", want, got)
	}
}
