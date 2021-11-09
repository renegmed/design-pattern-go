package main

func main(){
	// Initialise our sensors and stations
	coSensor := createWeatherSensor()
	wySensor:= createWeatherSensor()
	denverStation := createWeatherStation("Denver Weather Station")
	vailStation := createWeatherStation("Vail Weather Station")
	cheyenneStation := createWeatherStation("Cheyenne Weather Station")

	// Add the observer stations and notify them of current weather
	coSensor.addStation(denverStation)
	coSensor.addStation(vailStation)
	wySensor.addStation(cheyenneStation)
	coSensor.NotifyAll()
	wySensor.NotifyAll()

	// Change temperature for a new day
	coSensor.ChangeTemperature()
	wySensor.ChangeTemperature()

	// We can remove a station and maintain functionality
	coSensor.removeStation(vailStation)
	coSensor.ChangeTemperature()
	coSensor.NotifyAll()
	wySensor.NotifyAll()
}