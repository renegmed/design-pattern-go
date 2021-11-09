package main

import "fmt"

type BlueColour struct {
}

func (b *BlueColour) printColour() {
	fmt.Println("I'm Blue!")
}
