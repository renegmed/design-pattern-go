package pets

type Dog struct {
	Pet
}

func NewDog(name string, age int, sound string) *Dog {
	return &Dog{
		Pet: Pet{name, age, sound},
	}
}

// func (p *Dog) GetName() string {
// 	return p.name
// }

// func (p *Dog) GetSound() string {
// 	return p.sound
// }

// func (p *Dog) GetAge() int {
// 	return p.age
// }
