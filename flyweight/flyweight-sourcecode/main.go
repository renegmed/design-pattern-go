package main

// importing fmt package
import (
	"fmt"
)

//main method
func main() {
	//getFunc := func func(string, string) (DataDataTransferObject, error)

	var factory = NewDataTransferObjectFactory()
	dto, err := factory.getDataTransferObject("customer")
	if err != nil {
		panic(err)
	}
	fmt.Println("Customer ", dto.getId())

	dto, err = factory.getDataTransferObject("employee")
	if err != nil {
		panic(err)
	}
	fmt.Println("Employee ", dto.getId())

	dto, err = factory.getDataTransferObject("manager")
	if err != nil {
		panic(err)
	}
	fmt.Println("Manager ", dto.getId())

	dto, err = factory.getDataTransferObject("address")
	if err != nil {
		panic(err)
	}
	fmt.Println("Address ", dto.getId())

	// var customer DataTransferObject = factory.getDataTransferObject("customer", "1")
	// fmt.Println("Customer ", customer.getId())
	// fmt.Println("--------------------------")

	// var employee DataTransferObject = factory.getDataTransferObject("employee", "2")
	// fmt.Println("Employee ", employee.getId())

	// fmt.Println("--------------------------")
	// var manager DataTransferObject = factory.getDataTransferObject("manager", "2")
	// fmt.Println("Manager", manager.getId())

	// fmt.Println("--------------------------")
	// var address DataTransferObject = factory.getDataTransferObject("address", "3")
	// fmt.Println("Address", address.getId())
}
