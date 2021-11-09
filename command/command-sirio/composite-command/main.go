package main

import "fmt"

func main() {
	// ba := BankAccount{}
	// cmd := NewBankAccountCommand(&ba, Deposit, 100)
	// cmd.Call()
	// fmt.Println(ba)
	// cmd2 := NewBankAccountCommand(&ba, Withdraw, 25)
	// cmd2.Call()
	// fmt.Println(ba)
	// cmd2.Undo()
	// fmt.Println(ba)

	from := NewBankAccount("John", 0)
	cmd := NewBankAccountCommand(from, Deposit, 100)
	cmd.Call()
	fmt.Println(from)

	fmt.Println(".... Money Transfer ....")
	to := NewBankAccount("Peter", 0)
	mtc := NewMoneyTransferCommand(from, to, 25)

	mtc.Call()
	fmt.Println("From:", from)
	fmt.Println("To:", to)

	fmt.Println(".... Undo ....")
	mtc.Undo()
	fmt.Println("From:", from)
	fmt.Println("To:", to)
}
