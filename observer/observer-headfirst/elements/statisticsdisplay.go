package elements

import "fmt"

type statisticsDisplay struct {
	count   uint32
	avgTemp float32
	maxTemp float32
	minTemp float32
}

func NewStatisticsDisplay() *statisticsDisplay {
	return &statisticsDisplay{}
}

/**
 * This implements the avg, max and min temperatures.
 * When update is called, it does the calculation and
 * call the display() method to show the required information
 */
func (sd *statisticsDisplay) Update(temp float32, humidity float32, pressure float32) {
	sd.count++
	sd.avgTemp -= (sd.avgTemp - temp) / float32(sd.count)

	if sd.maxTemp < temp || sd.maxTemp == 0.0 {
		sd.maxTemp = temp
	}

	if sd.minTemp > temp || sd.minTemp == 0.0 {
		sd.minTemp = temp
	}

	sd.Display()
}

func (sd *statisticsDisplay) Display() {
	fmt.Printf("Statistics: Avg/Max/Min Temperature:%.2f, %.2f, %.2f\n", sd.avgTemp, sd.maxTemp, sd.minTemp)
}
