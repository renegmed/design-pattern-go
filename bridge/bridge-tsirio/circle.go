package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float32
	colour Colour // this is the bridge
}

func (c *Circle) getArea() float32 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) getPerimeter() float32 {
	return 2 * math.Pi * c.radius
}

func NewCircle(radius float32, colour Colour) *Circle {
	c := &Circle{radius: radius, colour: colour}
	return c
}

func (c *Circle) setColour(colour Colour) {
	c.colour = colour
}

func (c *Circle) getColour() {
	fmt.Print("Circle says: ")
	c.colour.printColour()
}
