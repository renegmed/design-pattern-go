package main

import "fmt"

//DataTransferObjectFactory struct
type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func NewDataTransferObjectFactory() *DataTransferObjectFactory {
	return &DataTransferObjectFactory{
		pool: make(map[string]DataTransferObject, 4),
	}
}

//DataTransferObjectFactory class method getDataTransferObject
func (factory *DataTransferObjectFactory) getDataTransferObject(dtoType string) (DataTransferObject, error) {

	dto, ok := factory.pool[dtoType]

	//if dto == nil {
	if ok {
		return dto, nil
	}

	fmt.Println("new DTO of dtoType: " + dtoType)
	switch dtoType {
	case "customer":
		factory.pool[dtoType] = NewCustomer("1") //{id: "1"}
	case "employee":
		factory.pool[dtoType] = NewEmployee("2") //{id: "2"}
	case "manager":
		factory.pool[dtoType] = NewManager("3") //{id: "3"}
	case "address":
		factory.pool[dtoType] = NewAddress("4") // {id: "4"}
	default:
		return nil, fmt.Errorf("Error: Cannot create %s object type", dtoType)
	}

	return factory.pool[dtoType], nil

}
