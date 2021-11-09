package builder

import "builder_design_pattern/builder/menu"

type director struct {
	builder iSubBuilder
}

func NewDirector() *director {
	return &director{}
}

func (d *director) SetBuilder(builder iSubBuilder) {
	d.builder = builder
}

func (d *director) BuildSub() menu.Sub {
	// d.builder.SetBread()
	// d.builder.SetCheese()
	// d.builder.SetToppings()
	// d.builder.SetSauces()

	return d.builder.GetSub()
}
