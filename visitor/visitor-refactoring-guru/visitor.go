package main

type visitor interface {
	visitForSquare(shape)
	visitForCircle(shape)
	visitForrectangle(shape)
}
