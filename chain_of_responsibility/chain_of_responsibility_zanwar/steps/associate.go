package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type associate struct {
	name string
	next step
}

func NewAssociate() *associate {
	return &associate{name: "associate"}
}

func (a *associate) Run(cust *customer.Customer) {
	if a.next == nil {
		fmt.Println("[Associate] Serving the regular customer: ", cust.Name)
		fmt.Println("[Associate] end of the line.")
		return
	}
	if cust.IsHighPriority {
		fmt.Println("Redirecting high priority customer directly to manager")
		fmt.Println("...next:", a.next.GetName())
		a.next.Run(cust)
		return
	}
	fmt.Println("[Associate] Serving the regular customer: ", cust.Name)
	fmt.Println("...next:", a.next.GetName())
	a.next.Run(cust)
}

func (a *associate) SetNextStep(next step) {
	a.next = next
}

func (a *associate) GetName() string {
	return a.name
}
