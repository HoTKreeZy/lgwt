package pae

type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposite(amout int) {
	w.balance += amout
}
