package main

import "fmt"

// CustomDateTemperatureService is another concrete subclass with a custom report format.
type CustomDateTemperatureService struct {
	BaseDateTemperatureService
	date        string
	temperature float64
}

// getDate returns the date for the CustomDateTemperatureService.
func (s *CustomDateTemperatureService) GetDate() string {
	return s.date
}

// getTemperature returns the temperature for the CustomDateTemperatureService.
func (s *CustomDateTemperatureService) GetTemperature() float64 {
	return s.temperature
}

// formatReport provides a custom implementation for the report format.
func (s *CustomDateTemperatureService) FormatReport(date string, temperature float64) string {
	return fmt.Sprintf("Custom Temperature Report: %s - %.1f°C", date, temperature)
}
