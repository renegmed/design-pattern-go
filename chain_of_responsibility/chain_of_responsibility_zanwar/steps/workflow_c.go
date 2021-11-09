package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type workflowC struct {
	name string
	next step
}

func NewWorkflowC() *workflowC {
	return &workflowC{name: "workflow C"}
}

func (w *workflowC) Run(cust *customer.Customer) {
	if w.next == nil {
		fmt.Println("[Workflow B] Serving the customer: ", cust.Name)
		fmt.Println("[Workflow B] end of the line.")
		return
	}
	fmt.Println("[Workflow B] Serving the customer: ", cust.Name)
	fmt.Println("...next:", w.next.GetName())
	w.next.Run(cust)
}

func (w *workflowC) SetNextStep(next step) {
	//fmt.Println("... workflow B SetNextStep")
	w.next = next
}
func (w *workflowC) GetName() string {
	return w.name
}
