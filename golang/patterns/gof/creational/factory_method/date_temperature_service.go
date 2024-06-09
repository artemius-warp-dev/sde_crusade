package main

import "fmt"

// DateTemperatureService struct with DateSensor and TemperatureSensor
type DateTemperatureService struct {
	date        *DateSensor
	temperature *TemperatureSensor
}

func (s *DateTemperatureService) ReportData() {
	fmt.Println(s.date.ReportData())
	fmt.Println(s.temperature.ReportData())
}
