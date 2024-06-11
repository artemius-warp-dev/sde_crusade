package main

type DateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
}

// SimpleDateTemperatureService is a basic implementation of DateTemperatureService.
type SimpleDateTemperatureService struct{}

// GetDate returns the current date.
func (s *SimpleDateTemperatureService) GetDate() string {
	return "2024-06-09"
}

// GetTemperature returns the current temperature.
func (s *SimpleDateTemperatureService) GetTemperature() float64 {
	return 25.5
}
