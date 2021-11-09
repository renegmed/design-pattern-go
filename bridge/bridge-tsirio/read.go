package main

import "fmt"

type RedColour struct{}

func (r *RedColour) printColour() {
	fmt.Println("I'm Red!")
}
