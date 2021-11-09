package main

import (
	"fmt"
	"math/rand"
)

const (
	maxTemperature = 120
	minTemperature = -30
)

type Subject interface{
    	NotifyAll()
}

type WeatherSensor struct{
	Subject
	weatherStations []*WeatherStation
	temperature int
}

func (ws *WeatherSensor) addStation(station *WeatherStation){
	ws.weatherStations = append(ws.weatherStations, station)
}

func (ws *WeatherSensor) removeStation(weatherStationToRemove *WeatherStation){
	for i, weatherStation := range ws.weatherStations {
		if weatherStation.Name == weatherStationToRemove.Name {
			/*
			Here we copy the last element of the array to the index of the station
			we want to delete. Then we truncate the last element of the array.
			*/
			ws.weatherStations[i] = ws.weatherStations[len(ws.weatherStations) - 1]
			ws.weatherStations = ws.weatherStations[:len(ws.weatherStations) - 1]
		}
	}
}

func (ws *WeatherSensor) NotifyAll(){
	for _, ob := range ws.weatherStations {
		ob.Update(ws.temperature)
	}
}

func createWeatherSensor() *WeatherSensor{
	return &WeatherSensor{
        temperature: getRandomTemperature(),
    }
}

func (ws *WeatherSensor) ChangeTemperature() {
    fmt.Printf("\nIt's a new day!\n")
    ws.temperature = getRandomTemperature()
    ws.NotifyAll()
}

func getRandomTemperature() int {
    return rand.Intn(maxTemperature - minTemperature) + minTemperature
}