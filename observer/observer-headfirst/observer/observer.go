package observer

/**
 * Observer interface is implemented by all observers,
 * Here we're passing the measurements to the observers
 * from the Subject when a weather measurement changes
 */
type Observer interface {
	Update(temp float32, humidity float32, pressure float32)
}
