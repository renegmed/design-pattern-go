package evictor

import (
	"fmt"
	"strategy_design_pattern/strategy"
)

type Lfu struct {
}

func (l *Lfu) Evict(c *strategy.Cache) {
	fmt.Println("Evicting by lfu strategy")
}
