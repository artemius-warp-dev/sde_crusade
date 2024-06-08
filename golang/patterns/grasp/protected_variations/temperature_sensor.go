package main

import "fmt"

// TemperatureSensor struct implementing the Sensor interface
type TemperatureSensor struct {
	temperature float64
}

// NewTemperatureSensor creates a new TemperatureSensor
func NewTemperatureSensor(temp float64) Sensor {
	return &TemperatureSensor{
		temperature: temp,
	}
}

// ReportData returns the temperature data
func (t *TemperatureSensor) ReportData() string {
	return fmt.Sprintf("Temperature: %.2f°C", t.temperature)
}
