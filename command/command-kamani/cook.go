package main

// A Cook comes with their list of commands as attributes
type Cook struct {
	Commands []Command
}

func (c *Cook) addCommand(command Command) {
	c.Commands = append(c.Commands, command)
}

// The executeCommands method executes all the commands
// one by one
func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}
