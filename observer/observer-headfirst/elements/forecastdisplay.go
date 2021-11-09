package elements

import "fmt"

type forecastDisplay struct {
	count   uint32
	avgTemp float32
	maxTemp float32
	minTemp float32
}

func NewForecastDisplay() *forecastDisplay {
	return &forecastDisplay{}
}

/**
 * This implements the avg, max and min temperatures.
 * When update is called, it does the calculation and
 * call the display() method to show the required information
 */
func (fd *forecastDisplay) Update(temp float32, humidity float32, pressure float32) {
	fd.count++
	fd.avgTemp -= (fd.avgTemp - temp) / float32(fd.count)

	if fd.maxTemp < temp || fd.maxTemp == 0.0 {
		fd.maxTemp = temp
	}

	if fd.minTemp > temp || fd.minTemp == 0.0 {
		fd.minTemp = temp
	}

	fd.Display()
}

func (fd *forecastDisplay) Display() {
	fmt.Printf("Forecast: Avg/Max/Min Temperature:%.2f, %.2f, %.2f\n", fd.avgTemp, fd.maxTemp, fd.minTemp)
}
