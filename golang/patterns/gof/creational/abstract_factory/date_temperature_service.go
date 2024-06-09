package main

import "fmt"

// DateTemperatureService struct with DateSensor and TemperatureSensor
type DateTemperatureService struct {
	dateSensor        DateSensor
	temperatureSensor TemperatureSensor
}

func (s *DateTemperatureService) ReportData() {
	fmt.Println(s.dateSensor.ReportData())
	fmt.Println(s.temperatureSensor.ReportData())
}
