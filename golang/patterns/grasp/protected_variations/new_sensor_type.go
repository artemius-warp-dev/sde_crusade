package main

import "fmt"

// NewSensorType struct implementing the Sensor interface
type NewSensorType struct {
	value float64
}

// NewNewSensorType creates a new NewSensorType
func NewNewSensorType(value float64) Sensor {
	return &NewSensorType{
		value: value,
	}
}

// ReportData returns the new sensor type data
func (n *NewSensorType) ReportData() string {
	return fmt.Sprintf("New Sensor Type Value: %.2f", n.value)
}
