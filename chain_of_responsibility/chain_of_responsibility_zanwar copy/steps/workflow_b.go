package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type workflowB struct {
	next step
}

func NewWorkflowB() *workflowB {
	return &workflowB{}
}

func (w *workflowB) Run(cust *customer.Customer) {
	fmt.Println("[Workflow B] Serving the customer: ", cust.Name)
	w.next.Run(cust)
}

func (w *workflowB) SetNextStep(next step) {
	//fmt.Println("... workflow B SetNextStep")
	w.next = next
}
