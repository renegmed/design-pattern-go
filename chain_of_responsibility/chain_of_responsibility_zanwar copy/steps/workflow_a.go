package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type workflowA struct {
	next step
}

func NewWorkflowA() *workflowA {
	return &workflowA{}
}

func (w *workflowA) Run(cust *customer.Customer) {
	fmt.Println("[Workflow A] Serving the customer: ", cust.Name)
	w.next.Run(cust)
}

func (w *workflowA) SetNextStep(next step) {
	w.next = next
}
