package main

import "fmt"

// DateSensor interface
type DateSensor interface {
	Sensor
	SetDate(date string)
}

// Concrete DateSensor implementation
type BasicDateSensor struct {
	date string
}

func (d *BasicDateSensor) ReportData() string {
	return fmt.Sprintf("Date: %s", d.date)
}

func (d *BasicDateSensor) SetDate(date string) {
	d.date = date
}
