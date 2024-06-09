// date_temperature_service.go
package main

// DateTemperatureService interface represents a service that provides date and temperature information.
type DateTemperatureService interface {
	DateSensor() string
	TemperatureSensor() float64
}

// NewDateTemperatureService represents the new service with different method names or signatures.
type NewDateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
}
