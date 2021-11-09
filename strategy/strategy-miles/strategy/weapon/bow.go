package weapon

import (
	"fmt"
	"strategy_design_pattern/strategy"
)

type bow struct {
	damage int
	name   string
}

func NewBow(name string, damage int) *bow {
	return &bow{
		name:   name,
		damage: damage,
	}
}

func (b *bow) UseWeapon(c *strategy.Character) {
	fmt.Printf("shoots the %s with a %s!\n", c.Name, b.name)
	c.Health -= b.damage
}
