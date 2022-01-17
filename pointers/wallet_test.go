package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, got, expect Bitcoin) {
		t.Helper()
		if got != expect {
			t.Errorf("expected %s but got %s", expect, got)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		expect := Bitcoin(10)
		assertBalance(t, got, expect)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		wallet.Withdraw(5)
		expect := Bitcoin(5)
		got := wallet.Balance()
		assertBalance(t, got, expect)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		expect := wallet.Withdraw(11)
		if expect == nil {
			t.Error("expected error but didn't get one")
		}
	})

}
