package main

import "fmt"

// TemperatureSensor struct implementing Sensor interface
type TemperatureSensor struct {
	temperature float64
}

func (t *TemperatureSensor) ReportData() string {
	return fmt.Sprintf("Temperature: %.2f°C", t.temperature)
}
