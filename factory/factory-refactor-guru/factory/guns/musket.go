package guns

// import (
// 	"factory_design_pattern/factory"
// )

type musket struct {
	gun
}

func NewMusket() *musket {
	return &musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
