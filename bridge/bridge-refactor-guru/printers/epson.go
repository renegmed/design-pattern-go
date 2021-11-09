package printers

import "fmt"

type epson struct {
}

func NewEpson() *epson {
	return &epson{}
}

func (p *epson) PrintFile() { // implements Printer interface
	fmt.Println("Printing by a EPSON Printer")
}
