package account

import "fmt"

//Account struct
type Account struct {
	id          string
	accountType string
}

func NewAccount() *Account {
	return &Account{}
}

//Account class method create - creates account given AccountType
func (account *Account) Create(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType

	return account
}

//Account class method getById  given id string
func (account *Account) GetById(id string) *Account {
	fmt.Println("getting account by Id")
	return account
}

func (account *Account) GetAccountType() string {
	return account.accountType
}

//Account class method deleteById given id string
func (account *Account) DeleteById(id string) {
	fmt.Println("delete account by id")
}
