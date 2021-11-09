package main

import (
	"fmt"
)

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s shape) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	square := s.(*square)
	area := square.getSide() * square.getSide()
	fmt.Println("Calculating area for", s.getType(), "  area:", area)
}

func (a *areaCalculator) visitForCircle(s shape) {
	c := s.(*circle)

	//var area float64

	area := float64(3.1415) * float64(c.getRadius()) * float64(c.getRadius())

	fmt.Printf("Calculating area for %s area: %.4f\n", s.getType(), area)
}
func (a *areaCalculator) visitForrectangle(s shape) {
	r := s.(*rectangle)
	area := r.getLength() * r.getWidth()
	fmt.Printf("Calculating area for %s  area: %d\n", s.getType(), area)
}
