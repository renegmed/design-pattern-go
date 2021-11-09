package main

import "fmt"

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account  *BankAccount
	action   Action
	amount   int
	succeded bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeded = true
	case Withdraw:
		b.succeded = b.account.Withdraw(b.amount)
	default:
		fmt.Println("...call() action:", b.action)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeded {
		return
	}
	// fmt.Println("... undo action:", b.action)
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	default:
		fmt.Println("...undo action invalid.")
	}
}
