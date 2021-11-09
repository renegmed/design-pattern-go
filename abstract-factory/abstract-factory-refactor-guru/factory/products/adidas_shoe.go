package products

type adidasShoe struct {
	shoe
}

func NewAdidasShoe() *adidasShoe {
	return &adidasShoe{}
}
