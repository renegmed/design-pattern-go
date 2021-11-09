package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type voiceAssistant struct {
	name string
	next step
}

func NewVoiceAssistant() *voiceAssistant {
	return &voiceAssistant{name: "voice assistant"}
}

func (v *voiceAssistant) Run(cust *customer.Customer) {
	if v.next == nil {
		fmt.Println("[Voice Assistant] Serving the customer: ", cust.Name)
		fmt.Println("[Voice Assistant] end of the line.")
		return
	}
	fmt.Println("[Voice Assistant] Serving the customer: ", cust.Name)
	fmt.Println("...next:", v.next.GetName())
	v.next.Run(cust)
}

func (v *voiceAssistant) SetNextStep(next step) {
	v.next = next
}

func (v *voiceAssistant) GetName() string {
	return v.name
}
