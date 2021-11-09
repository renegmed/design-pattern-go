package main

import "fmt"

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s shape) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c shape) {
	fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForrectangle(t shape) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
