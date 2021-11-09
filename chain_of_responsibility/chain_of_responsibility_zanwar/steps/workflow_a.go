package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type workflowA struct {
	name string
	next step
}

func NewWorkflowA() *workflowA {
	return &workflowA{name: "workflow A"}
}

func (w *workflowA) Run(cust *customer.Customer) {
	if w.next == nil {
		fmt.Println("[Workflow A] Serving the customer: ", cust.Name)
		fmt.Println("[Workflow A] end of the line.")
		return
	}
	fmt.Println("[Workflow A] Serving the customer: ", cust.Name)
	fmt.Println("...next:", w.next.GetName())
	w.next.Run(cust)
}

func (w *workflowA) SetNextStep(next step) {
	w.next = next
}

func (w *workflowA) GetName() string {
	return w.name
}
