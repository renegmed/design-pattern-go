package command

import (
	"command-design-pattern/receiver"
)

type ceilingFanOnCommand struct {
	ceilingFan *receiver.CeilingFan
}

func NewCeilingFanOnCommand(c *receiver.CeilingFan) *ceilingFanOnCommand {
	return &ceilingFanOnCommand{
		ceilingFan: c,
	}
}

func (l *ceilingFanOnCommand) Execute() {
	l.ceilingFan.On()
}

// ------------------

type ceilingFanOffCommand struct {
	ceilingFan *receiver.CeilingFan
}

func NewCeilingFanOffCommand(c *receiver.CeilingFan) *ceilingFanOffCommand {
	return &ceilingFanOffCommand{
		ceilingFan: c,
	}
}

func (l *ceilingFanOffCommand) Execute() {
	l.ceilingFan.Off()
}
