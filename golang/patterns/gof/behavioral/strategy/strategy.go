package main

// FormatStrategy interface defines a method for formatting date and temperature.
type FormatStrategy interface {
	Format(date string, temperature float64) string
}
