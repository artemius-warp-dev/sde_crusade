package main

func main() {

	// Create a new DateTemperatureService
	service := NewDateTemperatureService("2024-06-09", 25.5)
	service.Report()

	// Save the current state
	memento := service.Save()

	// Change the state
	service.SetState("2024-06-10", 30.0)
	service.Report()

	// Restore the saved state
	service.Restore(memento)
	service.Report()

}
