package main

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "Rectangle"
}

func (t *rectangle) getLength() int {
	return t.l
}

func (t *rectangle) getWidth() int {
	return t.b
}
