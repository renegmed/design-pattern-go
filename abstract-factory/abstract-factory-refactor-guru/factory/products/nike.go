package products

// import (
// 	"abstract_factory_design_factory/factory"
// )

type nike struct {
}

func NewNike() *nike {
	return &nike{}
}

func (n *nike) MakeShoe() IShoe { // implements iSportsFactory method
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) MakeShirt() IShirt { // implements iSportsFactory method
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 14,
		},
	}
}
