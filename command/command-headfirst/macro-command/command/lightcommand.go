package command

import (
	"command-design-pattern/receiver"
)

/**
 * The LightOffCommand works exactly
 * the same way as the LightOnCommand,
 * except that we are binding the receiver
 * to a different action: the off() method.
 */
type lightOffCommand struct {
	light *receiver.Light
}

func NewLightOffCommand(l *receiver.Light) *lightOffCommand {
	return &lightOffCommand{
		light: l,
	}
}

func (l *lightOffCommand) Execute() {
	l.light.Off()
}

// ------------------

type lightOnCommand struct {
	light *receiver.Light
}

func NewLightOnCommand(l *receiver.Light) *lightOnCommand {
	return &lightOnCommand{
		light: l,
	}
}

func (l *lightOnCommand) Execute() {
	l.light.On()
}
