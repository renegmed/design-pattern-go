package evictor

import (
	"fmt"
	"strategy_design_pattern/strategy"
)

type Lru struct {
}

func (l *Lru) Evict(c *strategy.Cache) {
	fmt.Println("Evicting by lru strategy")
}
