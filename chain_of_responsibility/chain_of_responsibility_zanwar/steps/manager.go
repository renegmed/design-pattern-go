package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type manager struct {
	name string
	next step
}

func NewManager() *manager {
	return &manager{name: "manager"}
}

func (m *manager) Run(cust *customer.Customer) {
	if m.next == nil {
		fmt.Println("[Manager] Serving the customer: ", cust.Name)
		fmt.Println("[Manager] end of the line")
		return
	}
	fmt.Println("[Manager] Serving the customer: ", cust.Name)
	fmt.Println("...next:", m.next.GetName())
}

func (m *manager) SetNextStep(next step) {
	m.next = next
}
func (m *manager) GetName() string {
	return m.name
}
