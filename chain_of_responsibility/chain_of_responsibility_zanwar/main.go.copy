package main

import (
	"chain_of_responsiblity_design_pattern/customer"
	"chain_of_responsiblity_design_pattern/steps"
	"fmt"
)

func main() {

	fmt.Println("===================")

	m := steps.NewManager() // &manager{}

	assoc := steps.NewAssociate() //&associate{}
	assoc.SetNextStep(m)

	va := steps.NewVoiceAssistant() //&voiceAssistant{}
	va.SetNextStep(assoc)

	// Chain formation complete

	// Start chain execution for normal customer
	// normalCust := &customer{
	// 	name: "Bob",
	// }

	normalCust := customer.NewCustomer("Bob", false)

	va.Run(normalCust)

	fmt.Println("===================")

	// Start chain execution for high priority customer
	// highPriorityCust := &customer{
	// 	name:           "John",
	// 	isHighPriority: true,
	// }

	highPriorityCust := customer.NewCustomer("John", true)

	va.Run(highPriorityCust)

	fmt.Println("===================")

	workflowB := steps.NewWorkflowB()
	workflowB.SetNextStep(m)

	workflowA := steps.NewWorkflowA()
	workflowA.SetNextStep(workflowB)

	va2 := steps.NewVoiceAssistant()
	va2.SetNextStep(workflowA)

	diffWorkflowCust := customer.NewCustomer("Peter", false)

	va2.Run(diffWorkflowCust)

}
