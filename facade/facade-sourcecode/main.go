package main

// importing fmt package
import (
	"facade_design_pattern/facade"
	"fmt"
)

//main method
func main() {
	var facade = facade.NewBranchManagerFacade()
	// var customer *customer.Customer
	// var account *account.Account

	customer, account := facade.CreateCustomerAccount("Thomas Smith", "Savings")
	fmt.Println(customer.GetName())
	fmt.Println(account.GetAccountType())
	var transaction = facade.CreateTransaction("21456", "87345", 1000)
	fmt.Println(transaction.GetAmount())

}
