package printers

import "fmt"

type hp struct {
}

func NewHp() *hp {
	return &hp{}
}

func (p *hp) PrintFile() { // implements Printer interface
	fmt.Println("Printing by a HP Printer")
}
