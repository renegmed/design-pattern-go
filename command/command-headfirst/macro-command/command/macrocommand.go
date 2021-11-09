package command

type macroCommand struct {
	commands []ICommand
}

func NewMacroCommand(c []ICommand) *macroCommand {
	return &macroCommand{
		commands: c,
	}
}

func (m *macroCommand) AddCommand(c ICommand) {
	m.commands = append(m.commands, c)
}

func (m *macroCommand) Execute() {
	for _, command := range m.commands {
		command.Execute()
	}
}
