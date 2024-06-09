package main

import "fmt"

func main() {
	// Get the singleton instance of DateTemperatureService
	service := GetInstance()
	service.ReportData()

	// Any subsequent calls to GetInstance will return the same instance
	anotherService := GetInstance()
	anotherService.ReportData()

	// Check if both instances are the same
	if service == anotherService {
		fmt.Println("Both instances are the same.")
	}
}
