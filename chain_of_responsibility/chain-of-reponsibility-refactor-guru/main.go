package main

import "fmt"

func main() {

	fmt.Println("+++++++++++++++++++++")

	patient := &patient{name: "abc"}

	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := newDoctor("John Gooding") //&doctor{}
	doctor.setNext(medical)

	doctor2 := newDoctor("Peter S. Shaw")
	doctor2.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	if patient.name == "abcx" {
		reception.setNext(doctor)
		doctor.setNext(medical)
	} else {
		reception.setNext(doctor2)
		doctor2.setNext(cashier)
	}

	//Patient visiting
	reception.execute(patient)

	patient.status()
}
