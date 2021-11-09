package facade

import (
	"facade_design_pattern/account"
	"facade_design_pattern/customer"
	"facade_design_pattern/transaction"
)

//BranchManagerFacade struct
type BranchManagerFacade struct {
	account     *account.Account
	customer    *customer.Customer
	transaction *transaction.Transaction
}

//methodd NewBranchManagerFacade
func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{
		account.NewAccount(),
		customer.NewCustomer(),
		transaction.NewTransaction(),
	} //{&Account{}, &Customer{}, &Transaction{}}
}

//BranchManagerFacade class method createCustomerAccount
func (facade *BranchManagerFacade) CreateCustomerAccount(customerName string, accountType string) (*customer.Customer, *account.Account) {
	var customer = facade.customer.Create(customerName)
	var account = facade.account.Create(accountType)
	return customer, account
}

//BranchManagerFacade class method createTransaction
func (facade *BranchManagerFacade) CreateTransaction(
	srcAccountId string, destAccountId string, amount float32) *transaction.Transaction {

	var transaction = facade.transaction.Create(srcAccountId, destAccountId, amount)
	return transaction

}
