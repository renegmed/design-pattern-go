package main

import (
	"fmt"
	"observer-design-pattern/elements"
)

func main() {
	/**
	 * Creating the weather object
	 */
	weatherData := elements.NewWeatherData()

	/**
	 * Creating the displays and registering as
	 * observer to the Subject
	 */
	currentConditionDisplay := elements.NewCurrentConditionDisplay()
	weatherData.RegisterObserver(currentConditionDisplay)

	statisticsDisplay := elements.NewStatisticsDisplay()
	weatherData.RegisterObserver(statisticsDisplay)
	forecastDisplay := elements.NewForecastDisplay()
	weatherData.RegisterObserver(forecastDisplay)

	/**
	 * Simulate new weather measurements
	 */
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(77, 90, 29.2)

	fmt.Println("---Unregister statistics display----")
	weatherData.DeregisterObserver(statisticsDisplay)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(77, 90, 29.2)

}
