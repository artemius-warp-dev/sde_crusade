package main

import "fmt"

// DateSensor struct implementing Sensor interface
type DateSensor struct {
	date string
}

func (d *DateSensor) ReportData() string {
	return fmt.Sprintf("Date: %s", d.date)
}
