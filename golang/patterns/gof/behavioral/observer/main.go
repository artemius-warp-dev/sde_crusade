package main

func main() {

	service := &SimpleDateTemperatureService{
		date:        "2024-06-09",
		temperature: 25.5,
	}

	// Create observers
	dateObserver := &DateObserver{}
	temperatureObserver := &TemperatureObserver{}

	// Add observers to the service
	service.AddObserver(dateObserver)
	service.AddObserver(temperatureObserver)

	// Change the date and temperature
	service.SetDate("2024-06-10")
	service.SetTemperature(26.0)

	// Reset the service
	service.Reset()

}
