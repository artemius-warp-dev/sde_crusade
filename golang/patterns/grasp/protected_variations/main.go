package main

func main() {
	// Create instances of different sensors
	tempSensor := NewTemperatureSensor(23.5)
	humiditySensor := NewHumiditySensor(45.7)
	newSensor := NewNewSensorType(78.9)

	// Create a BasicSensorManager
	manager := NewBasicSensorManager()

	// Create a SensorController with the manager
	controller := NewSensorController(manager)

	// Add sensors using the controller
	controller.AddSensor(tempSensor)
	controller.AddSensor(humiditySensor)
	controller.AddSensor(newSensor)

	// Report data from all sensors using the controller
	controller.ReportAllData()
}
