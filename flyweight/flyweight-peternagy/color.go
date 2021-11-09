package flyweight

import (
	"github.com/fatih/color"
)

//Color - struct holds color for reuse
type Color struct {
	Name  string
	Color *color.Color
}

//NewColor - create a new instance
func NewColor(name string) *Color {
	switch name {
	case "red":
		return &Color{Name: name, Color: color.New(color.FgRed)}
	case "blue":
		return &Color{Name: name, Color: color.New(color.FgBlue)}
	case "green":
		return &Color{Name: name, Color: color.New(color.FgGreen)}
	case "yellow":
		return &Color{Name: name, Color: color.New(color.FgYellow)}
	case "cyan":
		return &Color{Name: name, Color: color.New(color.FgCyan)}
	default:
		return &Color{Name: name, Color: color.New(color.FgWhite)}
	}
}

//PrintColoredln - print colored line ( to add some functionality)
func (c *Color) PrintColoredln(ln string) {
	c.Color.Println(ln)
}
