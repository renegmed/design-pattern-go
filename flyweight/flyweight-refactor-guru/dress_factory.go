package main

import "fmt"

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

type dressFactory struct {
	dressMap map[string]dress
}

var (
	dfsi *dressFactory = nil
)

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if _, ok := d.dressMap[dressType]; ok {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {

	if dfsi == (*dressFactory)(nil) {
		dfsi = &dressFactory{
			dressMap: make(map[string]dress),
		}
	}
	return dfsi

}
