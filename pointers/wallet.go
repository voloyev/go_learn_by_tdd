package goPointers

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#try-to-run-the-test-2

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int64

func (w *Wallet) Deposit(cash Bitcoin) {
	w.balance += cash
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(cash Bitcoin) error {
	if cash > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= cash

	return nil
}
