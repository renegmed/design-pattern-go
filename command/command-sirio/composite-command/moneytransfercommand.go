package main

type moneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *bankAccount
	amount   int
}

func NewMoneyTransferCommand(
	from *bankAccount, to *bankAccount, amount int) *moneyTransferCommand {

	c := &moneyTransferCommand{
		from:   from,
		to:     to,
		amount: amount,
	}

	c.commands = append(c.commands,
		NewBankAccountCommand(from, Withdraw, amount),
	)

	c.commands = append(c.commands,
		NewBankAccountCommand(to, Deposit, amount),
	)

	return c
}

func (m *moneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}
