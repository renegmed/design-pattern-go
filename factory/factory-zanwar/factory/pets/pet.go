package pets

// Pet is a struct that implements Pet interface and
// would be used in any animal struct that we create.
// See `Dog` and `Cat` below
type Pet struct {
	name  string
	age   int
	sound string
}

func (p *Pet) GetName() string {
	return p.name
}

func (p *Pet) GetSound() string {
	return p.sound
}

func (p *Pet) GetAge() int {
	return p.age
}
