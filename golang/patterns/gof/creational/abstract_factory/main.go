package main

func main() {
	// Create a factory instance
	var factory AbstractFactory = &BasicSensorFactory{}

	// Use the factory to create a DateTemperatureService
	service := factory.CreateDateTemperatureService()
	service.ReportData()

	// You can use the factory to create another service with different values
	anotherService := factory.CreateDateTemperatureService()
	anotherService.dateSensor.SetDate("2024-06-09")
	anotherService.temperatureSensor.SetTemperature(37.0)
	anotherService.ReportData()
}
