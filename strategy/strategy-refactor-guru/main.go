package main

import (
	"strategy_design_pattern/strategy"
	"strategy_design_pattern/strategy/evictor"
)

func main() {
	lfu := &evictor.Lfu{}
	cache := strategy.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &evictor.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &evictor.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}
