package birds

// ostrich extends bird struct
type ostrich struct {
	bird
}

func NewOstrich() *bird {
	o := &ostrich{}
	return o.bird.NewBird("Common Ostrich", 250, false)
	// return &ostrich{
	// 	bird: bird{
	// 		name:   "Common Ostrich",
	// 		weight: 250,
	// 		canFly: false,
	// 	},
	// }
}
