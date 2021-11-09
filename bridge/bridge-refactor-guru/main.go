package main

import (
	"bridge_design_pattern/computers"
	"bridge_design_pattern/printers"
	"fmt"
)

func main() {

	hpPrinter := printers.NewHp()       //&hp{}
	epsonPrinter := printers.NewEpson() // &epson{}

	macComputer := computers.NewMac() //&mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println("---------------")

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println("---------------")

	winComputer := computers.NewWindows() //&windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println("---------------")

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println("---------------")
}
