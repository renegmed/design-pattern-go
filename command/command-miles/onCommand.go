package main

type onCommand struct {
    device device
}

func newOnCommand(device device) *onCommand {
	return &onCommand{device}
}

func (c *onCommand) execute() {
    c.device.on()
}