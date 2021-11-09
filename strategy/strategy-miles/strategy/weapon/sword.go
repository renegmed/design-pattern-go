package weapon

import (
	"fmt"
	"strategy_design_pattern/strategy"
)

type sword struct {
	damage int
	name   string
}

func NewSword(name string, damage int) *sword {
	return &sword{
		name:   name,
		damage: damage,
	}
}

func (s *sword) UseWeapon(c *strategy.Character) {
	fmt.Printf("slashes the %s with a %s!\n", c.Name, s.name)
	c.Health -= s.damage
}
