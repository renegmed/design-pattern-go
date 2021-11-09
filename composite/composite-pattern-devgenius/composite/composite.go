package composite

import "fmt"

type Composite struct {
	component Component
	leaves    []Component
}

func (c *Composite) Add(leaf Component) {
	c.leaves = append(c.leaves, leaf)
}

func (c Composite) Display() {
	c.component.Display()
	if len(c.leaves) == 0 {
		return
	}
	fmt.Println("===List of Subordinates===")
	for _, leaf := range c.leaves {
		leaf.Display()
	}
	fmt.Println("===End===")
}

func NewComposite(component Component) Composite {
	return Composite{
		component: component,
	}
}
