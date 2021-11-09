// Author: Peter Nagy <https://peternagy.ie>
// Since: Dec, 2017
// Package: flyweight
// Description: flyweight patter implementation

package flyweight

import (
	"sync"
)

var (
	pool map[string]*Color
	once sync.Once
)

// func init() {
// 	once.Do(func() {
// 		pool = make(map[string]*Color)
// 	})
// }

//ColorFactory - get an instance from the initiaslized pool
func ColorFactory(name string) *Color {
	once.Do(func() {
		pool = make(map[string]*Color)
	})
	if v, ok := pool[name]; ok {
		return v
	}

	c := NewColor(name)
	pool[name] = c

	return c
}
