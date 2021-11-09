package elements

import (
	"observer-design-pattern/observer"
)

type weatherData struct {
	/**
	 * Implementation of set in Go
	 * Instantiating in the constructor (newWeatherData)
	 */
	observers   map[observer.Observer]bool
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() *weatherData {
	return &weatherData{
		observers: make(map[observer.Observer]bool),
	}
}

/**
 * WeatherData now implements the Subject interface
 */
func (w *weatherData) RegisterObserver(o observer.Observer) {
	/**
	 * When an observer registers, we add it in
	 * the map with value set as true
	 */
	w.observers[o] = true
}

func (w *weatherData) DeregisterObserver(o observer.Observer) {
	/**
	 * When an observer deregisters, we remove it
	 * from the map after checking if the value exists
	 */
	if _, ok := w.observers[o]; ok {
		delete(w.observers, o)
	}
}

func (w *weatherData) NotifyObservers() {
	/**
	 * This is where we tell all the observers about the state
	 * by calling the common update method
	 */
	for observer := range w.observers {
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

/**
 * We notify the Observers when we get updated
 * measurements from the Weather Station
 */
func (w *weatherData) measurementsChanged() {
	w.NotifyObservers()
}

/**
 * Dummy method to test our display elements.
 * Setting the measurements not via a device.
 */
func (w *weatherData) SetMeasurements(temp float32, humidity float32, pressure float32) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}
