package receiver

import "fmt"

type CeilingFan struct {
	Name string
}

func (c *CeilingFan) On() {
	fmt.Println(c.Name, "ceiling fan was turned on.")
}

func (c *CeilingFan) Off() {
	fmt.Println(c.Name, "ceiling fan was turned off.")
}
