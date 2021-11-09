package computers

import (
	"bridge_design_pattern/bridge"
	"fmt"
)

type windows struct {
	printer bridge.Printer
}

func NewWindows() *windows {
	return &windows{}
}
func (w *windows) Print() {
	fmt.Println("Print request for windows")
	w.printer.PrintFile()
}

func (w *windows) SetPrinter(p bridge.Printer) {
	w.printer = p
}
