package main

import "fmt"

func main() {
	// Create the initial state (normal state)
	initialState := &NormalState{
		date:        "2024-06-09",
		temperature: 25.5,
	}

	// Create the service context with the initial state
	service := NewDateTemperatureServiceContext(initialState)

	// Get the date and temperature in normal state
	fmt.Println("Date:", service.GetDate())
	fmt.Println("Temperature:", service.GetTemperature())

	// Transition to maintenance state
	maintenanceState := NewMaintenanceState(service)
	service.SetState(maintenanceState)

	// Try to get the date and temperature in maintenance state
	fmt.Println("Date (maintenance):", service.GetDate())
	fmt.Println("Temperature (maintenance):", service.GetTemperature())

	// Try to reset the service in maintenance state
	service.Reset()

	fmt.Println("Date (reset):", service.GetDate())
	fmt.Println("Temperature (reset):", service.GetTemperature())
}
