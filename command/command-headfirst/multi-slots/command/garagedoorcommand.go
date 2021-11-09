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

// -----------

type garageDoorCloseCommand struct {
	garage *receiver.Garage
}

func NewGarageDoorCloseCommand(g *receiver.Garage) *garageDoorCloseCommand {
	return &garageDoorCloseCommand{
		garage: g,
	}
}

func (g *garageDoorCloseCommand) Execute() {
	g.garage.Down()
	g.garage.LightOff()
}
