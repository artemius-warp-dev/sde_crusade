package main

type SimpleDateTemperatureImplementor struct{}

// FetchDate returns the current date.
func (s *SimpleDateTemperatureImplementor) FetchDate() string {
	return "2024-06-09"
}

// FetchTemperature returns the current temperature.
func (s *SimpleDateTemperatureImplementor) FetchTemperature() float64 {
	return 25.5
}
