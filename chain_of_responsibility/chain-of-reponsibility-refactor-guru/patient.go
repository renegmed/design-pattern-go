package main

import "fmt"

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func (p *patient) status() {
	fmt.Println("Patient name:", p.name)
	fmt.Println("\tRegistration is done?:", p.registrationDone)
	fmt.Println("\tDoctor checking is done?:", p.doctorCheckUpDone)
	fmt.Println("\tMedicine is done?:", p.medicineDone)
	fmt.Println("\tPayment is done?:", p.paymentDone)
}
