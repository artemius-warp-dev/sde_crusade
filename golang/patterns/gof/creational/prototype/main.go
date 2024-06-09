package main

import "fmt"

func main() {
	// Create the original DateTemperatureService instance
	originalService := &DateTemperatureService{
		date:        &DateSensor{date: "2024-06-08"},
		temperature: &TemperatureSensor{temperature: 36.6},
	}
	fmt.Println("Original")
	originalService.ReportData()

	// Clone the original service to create a copy
	clonedService := originalService.Clone()
	fmt.Println("Cloned")
	clonedService.ReportData()

	// Modify the cloned service to demonstrate that it is a separate instance
	clonedService.date.date = "2024-06-09"
	clonedService.temperature.temperature = 37.0
	fmt.Println("Cloned service is modified")

	// Report data from both original and cloned services to show they are independent
	fmt.Println("Original")
	originalService.ReportData()
	fmt.Println("Cloned modified")
	clonedService.ReportData()

}
