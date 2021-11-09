package main

import "fmt"

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 25)
	cmd2.Call()
	fmt.Println(ba)
	cmd2.Undo()
	fmt.Println(ba)
}
