package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type workflowB struct {
	name string
	next step
}

func NewWorkflowB() *workflowB {
	return &workflowB{name: "workflow B"}
}

func (w *workflowB) Run(cust *customer.Customer) {
	if w.next == nil {
		fmt.Println("[Workflow B] Serving the customer: ", cust.Name)
		fmt.Println("[Workflow B] end of the line.")
		return
	}
	fmt.Println("[Workflow B] Serving the customer: ", cust.Name)
	fmt.Println("...next:", w.next.GetName())
	w.next.Run(cust)
}

func (w *workflowB) SetNextStep(next step) {
	//fmt.Println("... workflow B SetNextStep")
	w.next = next
}
func (w *workflowB) GetName() string {
	return w.name
}
