package products

// import (
// 	"abstract_factory_design_factor/factory"
// )

type adidas struct {
}

func NewAdidas() *adidas {
	return &adidas{}
}

func (a *adidas) MakeShoe() IShoe { // implements iSportsFactory method
	as := &adidasShoe{}
	as.shoe = shoe{"adidas", 14}
	return as
	// return &adidasShoe{
	// 	shoe: shoe{
	// 		logo: "adidas",
	// 		size: 14,
	// 	},
	// }
}

func (a *adidas) MakeShirt() IShirt { // implements iSportsFactory method
	as := &adidasShirt{}
	as.shirt = shirt{"adidas", 14}
	return as

	// return &adidasShirt{
	// 	shirt: shirt{
	// 		logo: "adidas",
	// 		size: 14,
	// 	},
	// }
}
