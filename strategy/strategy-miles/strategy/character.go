package strategy

import (
	"fmt"
)

type Character struct {
	Health int
	Weapon weapon
	Name   string
	Damage int
}

func NewCharacter(name string) *Character {
	return &Character{
		Name:   name,
		Health: 100,
		Damage: 1,
	}
}

func (c *Character) EquipWeapon(w weapon) {
	c.Weapon = w
}

func (c *Character) Attack(opponent *Character) {
	fmt.Printf("The %s ", c.Name)
	c.Weapon.UseWeapon(opponent)
}
