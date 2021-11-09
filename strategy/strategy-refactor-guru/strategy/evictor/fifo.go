package evictor

import (
	"fmt"
	"strategy_design_pattern/strategy"
)

type Fifo struct {
}

func (l *Fifo) Evict(c *strategy.Cache) {
	fmt.Println("Evicting by fifo strategy")
}
