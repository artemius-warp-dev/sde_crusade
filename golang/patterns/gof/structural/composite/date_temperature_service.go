package main

// DateTemperatureService interface represents a service that provides date and temperature information.
type DateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
}
