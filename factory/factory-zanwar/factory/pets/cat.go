package pets

type Cat struct {
	Pet
}

func NewCat(name string, age int, sound string) *Cat {
	return &Cat{
		Pet: Pet{name, age, sound},
	}
}

// func (p *Cat) GetName() string {
// 	return p.name
// }

// func (p *Cat) GetSound() string {
// 	return p.sound
// }

// func (p *Cat) GetAge() int {
// 	return p.age
// }
