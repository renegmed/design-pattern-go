package main

import (
	"chain_of_responsiblity_design_pattern/customer"
	"chain_of_responsiblity_design_pattern/steps"
	"fmt"
	"os"
)

const HIGH_PRIORITY = true
const LOW_PRIORITY = false

func main() {
	var choice int
	for {
		fmt.Println("Which flow to use?")
		fmt.Print("Enter 1=normal 2=high piority 3=Peter's flow 4=other's flow  5=exit:")

		n, err := fmt.Scanf("%d", &choice)
		if n != 1 || err != nil {
			fmt.Println("Follow directions!")
			return
		}
		switch choice {
		case 1:
			normal()
		case 2:
			highPriority()
		case 3:
			differentFlow("Peter")
		case 4:
			differentFlow("John")
		case 5:
			os.Exit(2)
		}
	}

}

func normal() {

	fmt.Println("===================")

	m := steps.NewManager() // &manager{}

	assoc := steps.NewAssociate() //&associate{}
	assoc.SetNextStep(m)

	va := steps.NewVoiceAssistant() //&voiceAssistant{}
	va.SetNextStep(assoc)

	normalCust := customer.NewCustomer("Bob", LOW_PRIORITY)

	va.Run(normalCust) // starting from voice assistant

}

func highPriority() {

	fmt.Println("===================")

	// from voice assistant to manager by-passing associate
	manager := steps.NewManager()
	assoc := steps.NewAssociate()
	va := steps.NewVoiceAssistant()

	assoc.SetNextStep(manager)
	va.SetNextStep(assoc)

	highPriorityCust := customer.NewCustomer("John", HIGH_PRIORITY)

	va.Run(highPriorityCust) // starting from voice assistant
}

func differentFlow(name string) {
	fmt.Println("===================")

	diffWorkflowCust := customer.NewCustomer(name, LOW_PRIORITY)

	manager := steps.NewManager()
	workflowA := steps.NewWorkflowA()
	workflowB := steps.NewWorkflowB()
	workflowC := steps.NewWorkflowC()
	va := steps.NewVoiceAssistant()

	// start with voice assistant
	// then to workflow A
	// then to workflow B
	// then to manager
	// Exception: if customer name is Peter, by-pass worflow B
	va.SetNextStep(workflowA)

	if diffWorkflowCust.GetName() == "Peter" {
		workflowA.SetNextStep(workflowC)
		workflowC.SetNextStep(workflowB)
		workflowB.SetNextStep(manager)
	} else {
		workflowA.SetNextStep(workflowB)
		workflowB.SetNextStep(manager)
	}

	va.Run(diffWorkflowCust) // starting from voice assistant
}
