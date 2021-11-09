package main

import (
	"fmt"
)

type Runner struct {
	run func()
}

func (r *Runner) SetRun(f func()) {
	r.run = f
}

func (r *Runner) Start() {
	// some prepare stuff...
	if r.run == nil {
		fmt.Println("No function set.")
		return
	}
	r.run()
}

func runLog() {
	fmt.Println("Running")
}

func runLog2() {
	fmt.Println("Running logger 2")
}

func runLog3() {
	fmt.Println("Running logger 3")
}
func NewLogger(f func()) *Runner {
	return &Runner{run: f}
}

func main() {
	l := NewLogger(runLog) // pass a function
	l.Start()

	l2 := NewLogger(runLog2)
	l2.Start()

	l3 := NewLogger(nil)
	l3.Start()
	l3.SetRun(runLog3)
	l3.Start()
}
