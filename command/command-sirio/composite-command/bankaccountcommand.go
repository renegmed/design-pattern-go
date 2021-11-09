package main

import "fmt"

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type bankAccountCommand struct {
	account   *bankAccount
	action    Action
	amount    int
	succeeded bool
}

func NewBankAccountCommand(
	account *bankAccount, action Action, amount int) *bankAccountCommand {
	return &bankAccountCommand{
		account: account,
		action:  action,
		amount:  amount,
	}
}

func (b *bankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	default:
		fmt.Println("...call() action:", b.action)
	}
}

func (b *bankAccountCommand) Undo() {
	if !b.succeeded {
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

func (b *bankAccountCommand) Succeeded() bool {
	return b.succeeded
}

func (b *bankAccountCommand) SetSucceeded(value bool) {
	b.succeeded = value
}
