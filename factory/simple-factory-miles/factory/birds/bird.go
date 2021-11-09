package birds

import "fmt"

type bird struct {
	name   string
	weight float64
	canFly bool
}

func (b *bird) NewBird(name string, weight float64, canFly bool) *bird {
	return &bird{name, weight, canFly}
}

func (b *bird) SetName(name string) {
	b.name = name
}

func (b *bird) GetName() string {
	return b.name
}

func (b *bird) TryToFly() {
	if b.canFly {
		fmt.Printf("The %s is flying!\n", b.name)
		return
	}
	fmt.Printf("The %s cannot fly.\n", b.name)
	return
}
