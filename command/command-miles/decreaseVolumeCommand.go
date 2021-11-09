package main

type decreaseVolumeCommand struct {
    device device
}

func newDecreaseVolumeCommand(device device) *decreaseVolumeCommand {
	return &decreaseVolumeCommand{device}
}

func (c *decreaseVolumeCommand) execute() {
    c.device.decreaseVolume()
}