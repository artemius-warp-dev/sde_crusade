package main

import "fmt"

// DetailedFormatStrategy formats the date and temperature in a detailed way.
type DetailedFormatStrategy struct{}

// Format returns a detailed formatted string with date and temperature.
func (d *DetailedFormatStrategy) Format(date string, temperature float64) string {
	return fmt.Sprintf("Detailed Report - Date: %s, Temperature: %.1f degrees Celsius", date, temperature)
}
