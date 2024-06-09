package main

// DateTemperatureService interface represents a service that provides date and temperature information.
type DateTemperatureService interface {
	DateSensor() string
	TemperatureSensor() float64
}
