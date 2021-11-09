package shape

import "fmt"

type Draw3DShape struct{}

func NewDraw3DShape() IDrawShape {
	return Draw3DShape{}
}

// DrawShape struct has  method draw Shape with float x and y coordinates
func (d Draw3DShape) DrawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing 3D Shape:\n\tx:", x, "\n\ty:", y)
}
