package bridge

import (
	"bridge_design_pattern/printers"
)

type Computer interface {
	Print()
	SetPrinter(printers.Printer)
}
