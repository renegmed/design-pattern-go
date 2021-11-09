package command

import (
	"command-design-pattern/receiver"
)

type garageDoorOpenCommand struct {
	garage *receiver.Garage
}

func NewGarageDoorOpenCommand(g *receiver.Garage) *garageDoorOpenCommand {
	return &garageDoorOpenCommand{
		garage: g,
	}
}

func (g *garageDoorOpenCommand) Execute() {
	g.garage.Up()
	g.garage.LightOn()
}
