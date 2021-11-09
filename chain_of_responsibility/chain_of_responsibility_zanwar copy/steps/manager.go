package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type manager struct {
	next step
}

func NewManager() *manager {
	return &manager{}
}

func (m *manager) Run(cust *customer.Customer) {
	fmt.Println("[Manager] Serving the customer: ", cust.Name)
}

func (m *manager) SetNextStep(next step) {
	m.next = next
}
