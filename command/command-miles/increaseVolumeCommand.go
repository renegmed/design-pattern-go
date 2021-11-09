package main

type increaseVolumeCommand struct {
    device device
}

func newIncreaseVolumeCommand(device device) *increaseVolumeCommand {
	return &increaseVolumeCommand{device}
}
func (c *increaseVolumeCommand) execute() {
    c.device.increaseVolume()
}