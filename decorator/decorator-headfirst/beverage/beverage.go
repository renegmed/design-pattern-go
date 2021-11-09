package beverage

/**
 * For our case, we can use interface for beverage
 * instead of an abstract class.
 */
type Beverage interface {
	Description() string
	Cost() float32
}
