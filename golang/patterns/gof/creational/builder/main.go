package main

func main() {
	// Create a new builder instance
	builder := &DateTemperatureServiceBuilder{}

	// Use the builder to set the date and temperature and build the service
	service := builder.SetDate("2024-06-08").SetTemperature(36.6).Build()
	service.ReportData()

	// You can reuse the builder to create another service with different values
	anotherService := builder.SetDate("2024-06-09").SetTemperature(37.0).Build()
	anotherService.ReportData()
}
