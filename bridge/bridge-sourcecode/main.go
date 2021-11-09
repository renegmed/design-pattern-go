package main

import (
	"bridge_design_pattern/contour"
	"bridge_design_pattern/shape"
)

// main method
func main() {

	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}

	//var contour contour.IContour = contour.DrawContour{x, y, shape.DrawShape{}, 2} // DrawContour{} is the bridge

	//var contour contour.IContour = contour.DrawContour{x, y, shape.NewDrawShape(), 2} // DrawContour{} is the bridge
	var ctr = contour.NewDrawContour(x, y, shape.NewDrawShape(), 2)
	ctr.ResizeByFactor(2)
	ctr.DrawContour(x, y)

	ctr = contour.NewDrawContour(x, y, shape.NewDraw3DShape(), 2)
	ctr.ResizeByFactor(3)
	ctr.DrawContour(x, y)

}
