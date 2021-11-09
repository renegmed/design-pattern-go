package transaction

import "fmt"

//Transaction struct
type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

//Transaction class method create Transaction
func (transaction *Transaction) Create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

func (transaction *Transaction) GetAmount() float32 {
	return transaction.amount
}
