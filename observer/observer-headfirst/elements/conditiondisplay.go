package elements

import "fmt"

type currentConditionDisplay struct {
	temperature float32
	humidity    float32
	pressure    float32
}

func NewCurrentConditionDisplay() *currentConditionDisplay {
	return &currentConditionDisplay{}
}

/**
 * This implements Observer and DisplayElement interfaces to get changes from WeatherData
 * and show the information based on its functionality respectively
 */
func (ccd *currentConditionDisplay) Update(temp float32, humidity float32, pressure float32) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.pressure = pressure
	/**
	 * When update() is called, we save the measurements
	 * and call display() to show the required information
	 */
	ccd.Display()
}

func (ccd *currentConditionDisplay) Display() {
	fmt.Printf("Current conditions: Temperature:%.2f, Humidity:%.2f and Pressure:%.2f\n", ccd.temperature, ccd.humidity, ccd.pressure)
}
