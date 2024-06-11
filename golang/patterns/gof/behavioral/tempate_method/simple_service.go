package main

import "fmt"

// SimpleDateTemperatureService is a concrete subclass that implements the required methods.
type SimpleDateTemperatureService struct {
	BaseDateTemperatureService
	date        string
	temperature float64
}

// getDate returns the date for the SimpleDateTemperatureService.
func (s *SimpleDateTemperatureService) GetDate() string {
	return s.date
}

// getTemperature returns the temperature for the SimpleDateTemperatureService.
func (s *SimpleDateTemperatureService) GetTemperature() float64 {
	return s.temperature
}

// formatReport provides a custom implementation for the report format.
func (s *SimpleDateTemperatureService) FormatReport(date string, temperature float64) string {
	return fmt.Sprintf("Simple temperature Report: %s - %.1f°C", date, temperature)
}
