package main

import "fmt"

type player struct {
	name       string
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(name, playerType, dressType string) (*player, error) {
	dress, err := getDressFactorySingleInstance().getDressByType(dressType)
	if err != nil {
		return nil, err
	}

	return &player{
		name:       name,
		playerType: playerType,
		dress:      dress,
	}, nil
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

func (p *player) describe() {
	fmt.Printf("Player:\n\tname: %s\n\tplayer type:%s\n\tdress color:%s\n", p.name, p.playerType, p.dress.getColor())

}
