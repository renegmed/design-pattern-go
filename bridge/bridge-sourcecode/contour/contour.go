package contour

import (
	"bridge_design_pattern/shape"
	"fmt"
)

//DrawContour struct
type DrawContour struct {
	X      [5]float32
	Y      [5]float32
	Shape  shape.IDrawShape //shape.DrawShape
	Factor int
}

func NewDrawContour(x, y [5]float32, drawShape shape.IDrawShape, factor int) IContour {
	return DrawContour{X: x, Y: y, Shape: drawShape, Factor: factor}
}

//DrawContour method drawContour given the coordinates
func (contour DrawContour) DrawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.Shape.DrawShape(contour.X, contour.Y) // uses the bridge shape(DrawShape{} object)
}

//DrawContour method resizeByFactor given factor
func (contour DrawContour) ResizeByFactor(factor int) {
	contour.Factor = factor
}
