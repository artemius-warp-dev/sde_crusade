package main

import "fmt"

func main() {
	// Create the factory
	var factory ServiceFactory = &DateTemperatureServiceFactory{}

	// Use the factory to create a DateTemperatureService
	service := factory.CreateService()
	service.ReportData()

	// You can create multiple services using the factory
	anotherService := factory.CreateService()
	anotherService.ReportData()

	if service != anotherService {
		fmt.Println("Both services are not the same")
	}
}
