package main

import "fmt"

// HumiditySensor struct implementing the Sensor interface
type HumiditySensor struct {
	humidity float64
}

// NewHumiditySensor creates a new HumiditySensor
func NewHumiditySensor(humidity float64) Sensor {
	return &HumiditySensor{
		humidity: humidity,
	}
}

// ReportData returns the humidity data
func (h *HumiditySensor) ReportData() string {
	return fmt.Sprintf("Humidity: %.2f%%", h.humidity)
}
