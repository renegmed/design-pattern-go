package factory

import (
	"abstract_factory_design_factory/factory/products"
	"fmt"
)

type iSportsFactory interface {
	MakeShoe() products.IShoe
	MakeShirt() products.IShirt
}

func GetSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return products.NewAdidas(), nil //&adidas{}, nil
	}

	if brand == "nike" {
		return products.NewNike(), nil //&nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
