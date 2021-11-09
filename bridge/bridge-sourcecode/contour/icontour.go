package contour

//IContour interace
type IContour interface {
	DrawContour(x [5]float32, y [5]float32)
	ResizeByFactor(factor int)
}
