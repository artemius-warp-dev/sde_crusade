package main

import "fmt"

// TemperatureSensor interface
type TemperatureSensor interface {
	Sensor
	SetTemperature(temp float64)
}

// Concrete TemperatureSensor implementation
type BasicTemperatureSensor struct {
	temperature float64
}

func (t *BasicTemperatureSensor) ReportData() string {
	return fmt.Sprintf("Temperature: %.2f°C", t.temperature)
}

func (t *BasicTemperatureSensor) SetTemperature(temp float64) {
	t.temperature = temp
}
