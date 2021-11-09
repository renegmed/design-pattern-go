package main

type CompositeBankAccountCommand struct {
	commands []Command
}

// func (c *CompositeBankAccountCommand) Call() {
// 	for i := range c.commands {
// 		c.commands[i].Call()
// 	}
// }

// the Undo method will cycle backwards through all
// the commands and Undo them
func (c *CompositeBankAccountCommand) Undo() {
	for idx := range c.commands {
		c.commands[len(c.commands)-idx-1].Undo()
	}
}

// The Succeeded Getter will ask if there's at least
// one failed command and return false, otherwise
// everything is Ok
func (c *CompositeBankAccountCommand) Succeeded() bool {
	for idx := range c.commands {
		if c.commands[idx].Succeeded() {
			return false
		}
	}
	return true
}

// The Succeeded Setter will set the succeeded value with
// the operations status
func (c *CompositeBankAccountCommand) SetSucceeded(value bool) {
	for idx := range c.commands {
		c.commands[idx].SetSucceeded(value)
	}
}
