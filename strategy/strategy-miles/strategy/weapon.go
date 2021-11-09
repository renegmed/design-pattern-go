package strategy

type weapon interface {
	UseWeapon(opponent *Character)
}
