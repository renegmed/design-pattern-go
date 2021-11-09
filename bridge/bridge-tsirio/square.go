package main

import "fmt"

type Square struct {
	side   float32
	colour Colour // this is the bridge
}

func (s *Square) getArea() float32 {
	return s.side * s.side
}

func (s *Square) getPerimeter() float32 {
	return s.side * 4
}

func NewSquare(side float32, colour Colour) *Square {
	s := &Square{side: side, colour: colour}
	return s
}

func (s *Square) setColour(colour Colour) {
	s.colour = colour
}

func (s *Square) getColour() {
	fmt.Print("Square says: ")
	s.colour.printColour()
}
