package shape

import "fmt"

type DrawShape struct{}

func NewDrawShape() IDrawShape {
	return DrawShape{}
}

// DrawShape struct has  method draw Shape with float x and y coordinates
func (d DrawShape) DrawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape:\n\tx:", x, "\n\ty:", y)
}
