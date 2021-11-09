package computers

import (
	"bridge_design_pattern/bridge"
	"fmt"
)

type mac struct {
	printer bridge.Printer
}

func NewMac() *mac {
	return &mac{}
}

func (m *mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile() // Printer interface method
}

func (m *mac) SetPrinter(p bridge.Printer) {
	m.printer = p
}
