package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type associate struct {
	next step
}

func NewAssociate() *associate {
	return &associate{}
}

func (a *associate) Run(cust *customer.Customer) {
	if cust.IsHighPriority {
		fmt.Println("Redirecting high priority customer directly to manager")
		a.next.Run(cust)
		return
	}
	fmt.Println("[Associate] Serving the regular customer: ", cust.Name)
	a.next.Run(cust)
}

func (a *associate) SetNextStep(next step) {
	a.next = next
}
