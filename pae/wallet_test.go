package pae

import "testing"

func TestWaller(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposite(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
