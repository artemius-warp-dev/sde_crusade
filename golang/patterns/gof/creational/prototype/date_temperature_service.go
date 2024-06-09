package main

import "fmt"

// DateTemperatureService struct with DateSensor and TemperatureSensor
type DateTemperatureService struct {
	date        *DateSensor
	temperature *TemperatureSensor
}

// Clone method for DateTemperatureService to create a copy of the service
func (s *DateTemperatureService) Clone() *DateTemperatureService {
	return &DateTemperatureService{
		date:        s.date.Clone(),
		temperature: s.temperature.Clone(),
	}
}

// ReportData reports data from both DateSensor and TemperatureSensor
func (s *DateTemperatureService) ReportData() {
	fmt.Println(s.date.ReportData())
	fmt.Println(s.temperature.ReportData())
}
