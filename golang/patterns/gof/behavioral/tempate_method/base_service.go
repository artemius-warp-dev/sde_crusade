package main

import "fmt"

// DateTemperatureService defines the template method for reporting date and temperature.
type DateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
	FormatReport(date string, temperature float64) string
}

// BaseDateTemperatureService provides a default implementation for the template method.
type BaseDateTemperatureService struct {
	service DateTemperatureService
}

// Report is the template method that defines the algorithm's structure.
func (s *BaseDateTemperatureService) Report() {
	date := s.service.GetDate()
	temperature := s.service.GetTemperature()
	report := s.service.FormatReport(date, temperature)
	fmt.Println(report)
}
