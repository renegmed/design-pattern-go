package main

import "fmt"

func main() {
	r := &RedColour{}
	b := &BlueColour{}

	s := NewSquare(4.2, r)
	c := NewCircle(6.9, b)

	fmt.Println("Square's area is:", s.getArea())
	fmt.Println("Square's perimeter is:", s.getPerimeter())
	s.getColour()

	s.setColour(b)
	s.getColour()

	fmt.Println()

	fmt.Println("Circle's area is:", c.getArea())
	fmt.Println("Circle's perimeter is:", c.getPerimeter())
	c.getColour()

	c.setColour(r)
	c.getColour()
}
