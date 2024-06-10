package main

import "fmt"

// SimpleFormatStrategy formats the date and temperature in a simple way.
type SimpleFormatStrategy struct{}

// Format returns a simple formatted string with date and temperature.
func (s *SimpleFormatStrategy) Format(date string, temperature float64) string {
	return "Date: " + date + ", Temperature: " + fmt.Sprintf("%.1f", temperature)
}
