package main

import (
	"fmt"
	"strategy_design_pattern/strategy"
	"strategy_design_pattern/strategy/weapon"
)

func printCharacterStats(c *strategy.Character) {
	fmt.Printf("The %s has %d health left.\n", c.Name, c.Health)
}

func main() {

	// ----- setup the the weapons to be used
	godSword := weapon.NewSword("Armadyl Godsword", 45)
	darkBow := weapon.NewBow("Dark Bow", 35)
	giantSword := weapon.NewSword("Giant's Sword", 55)

	// ------ create a character and arm with a weapon

	champion := strategy.NewCharacter("Champion")
	champion.EquipWeapon(godSword)

	// ------ create another character and arm with weapon
	troll := strategy.NewCharacter("Cave Troll")
	troll.EquipWeapon(giantSword)

	printCharacterStats(champion)
	printCharacterStats(troll)

	// ---- start the battle, changing weapons ------

	champion.Attack(troll)

	printCharacterStats(troll)

	troll.Attack(champion)

	printCharacterStats(champion)

	champion.EquipWeapon(darkBow)
	champion.Attack(troll)

	printCharacterStats(troll)
}
