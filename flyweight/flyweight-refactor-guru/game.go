package main

type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		// terrorists:        make([]*player, 1),
		// counterTerrorists: make([]*player, 1),
	}
}

func (c *game) addTerrorist(name, dressType string) error {
	p, err := newPlayer(name, "T", dressType)
	if err != nil {
		return err
	}

	c.terrorists = append(c.terrorists, p)

	return nil

}

func (c *game) addCounterTerrorist(name, dressType string) error {
	p, err := newPlayer(name, "CT", dressType)
	if err != nil {
		return err
	}
	c.counterTerrorists = append(c.counterTerrorists, p)

	return nil
}

func (c *game) getTerrorists() []*player {

	return c.terrorists
}

func (c *game) getCounterTerrorists() []*player {
	return c.counterTerrorists
}
