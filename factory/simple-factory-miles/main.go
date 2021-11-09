package main

import (
	"factory_design_pattern/factory"
	"fmt"
)

func main() {
	ostrich, _ := factory.MakeBird("ostrich")
	osprey, _ := factory.MakeBird("osprey")
	_, err := factory.MakeBird("dodo")
	if err != nil {
		fmt.Println(err)
	}

	ostrich.TryToFly()
	osprey.TryToFly()
}
