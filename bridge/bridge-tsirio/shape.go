package main

type Shape interface {
	getArea() float32
	getPerimeter() float32
	getColour()
	setColour(colour Colour)
}
