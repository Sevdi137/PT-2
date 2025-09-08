package main

import (
	"fmt"
	"time"
)

type SensorData struct {
	SensorID   string
	Temperature float64
	Humidity    float64
	Timestamp   time.Time
}

var sensorDataList []SensorData

func main() {
	sensorDataList = append(sensorDataList, SensorData{"Датчик1", 22.5, 60.0, time.Now()})
	sensorDataList = append(sensorDataList, SensorData{"Датчик2", 23.0, 55.0, time.Now()})
	sensorDataList = append(sensorDataList, SensorData{"Датчик3", 21.0, 65.0, time.Now()})
	sensorDataList = append(sensorDataList, SensorData{"Датчик4", 24.0, 50.0, time.Now()})
	sensorDataList = append(sensorDataList, SensorData{"Датчик5", 20.0, 70.0, time.Now()})
	fmt.Println("Датчики:", sensorDataList)
	avgTemp := avgTemperature()
	fmt.Printf("Средняя температура: %.2f°C\n", avgTemp)
}

func avgTemperature() float64 {
	var total float64
	for _, data := range sensorDataList {
		total += data.Temperature
	}
	return total / float64(len(sensorDataList))
}

