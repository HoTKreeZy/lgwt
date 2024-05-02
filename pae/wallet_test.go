package pae

import "testing"

func TestWaller(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposite(Bitcoin(10))

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
