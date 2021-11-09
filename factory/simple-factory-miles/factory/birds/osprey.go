package birds

// osprey extends bird struct
type osprey struct {
	bird
}

func NewOsprey() *bird {
	b := &osprey{}
	return b.bird.NewBird("Western Osprey", 3.25, true)
	// return &osprey{
	// 	bird: factory.Bird{
	// 		name:   "Western Osprey",
	// 		weight: 3.25,
	// 		canFly: true,
	// 	},

}
