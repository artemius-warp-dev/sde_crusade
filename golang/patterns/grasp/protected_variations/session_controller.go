package main

// SensorController struct controlling the sensor operations
type SensorController struct {
	manager SensorManager
}

// NewSensorController creates a new SensorController
func NewSensorController(manager SensorManager) *SensorController {
	return &SensorController{
		manager: manager,
	}
}

// AddSensor adds a sensor using the manager
func (sc *SensorController) AddSensor(sensor Sensor) {
	sc.manager.AddSensor(sensor)
}

// ReportAllData reports all data using the manager
func (sc *SensorController) ReportAllData() {
	sc.manager.ReportAllData()
}
