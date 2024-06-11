package main

// DateTemperatureService interface represents a service that provides date and temperature information.
type DateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
	Reset()
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

// Reset resets the service.
func (s *SimpleDateTemperatureService) Reset() {
	// For demonstration, we'll just print a message.
	// In a real implementation, you might reset internal state here.
	println("Service has been reset.")
}
