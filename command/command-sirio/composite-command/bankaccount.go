package main

import "fmt"

var overdraftLimit = -500

type bankAccount struct {
	name    string
	balance int
}

func NewBankAccount(name string, balance int) *bankAccount {
	return &bankAccount{name, balance}
}
func (b *bankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Printf("Deposited %s account: %d, balance is now %d\n", b.name, amount, b.balance)
}

func (b *bankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Printf("Withdrew %s account: %d, balance is now %d\n", b.name, amount, b.balance)
		return true
	}
	return false
}
