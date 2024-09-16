package accounts

import (
	"errors"
	"fmt"
)

// Account
type Account struct {
	owner	string
	balance int
}

// function
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// method
// Deposit. receiver이 this 받아서 extension method 처럼 처리하는 거구나.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance
func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("can't widthdraw")

// Withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}