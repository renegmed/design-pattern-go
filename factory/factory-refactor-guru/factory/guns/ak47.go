package guns

// import (
// 	"factory_design_pattern/factory"
// )

type ak47 struct {
	gun
}

func NewAk47() *ak47 {
	// a := ak47{}
	// return a.gun.NewGun("AK47 gun", 4)
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
