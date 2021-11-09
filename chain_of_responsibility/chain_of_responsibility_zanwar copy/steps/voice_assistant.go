package steps

import (
	"chain_of_responsiblity_design_pattern/customer"
	"fmt"
)

type voiceAssistant struct {
	next step
}

func NewVoiceAssistant() *voiceAssistant {
	return &voiceAssistant{}
}

func (v *voiceAssistant) Run(cust *customer.Customer) {
	fmt.Println("[Voice Assistant] Serving the customer: ", cust.Name)
	v.next.Run(cust)
}

func (v *voiceAssistant) SetNextStep(next step) {
	v.next = next
}
