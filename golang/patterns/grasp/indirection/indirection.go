package main

import "fmt"

// Diffirent sensor types

// Sensor interface with a ReportData method
type Sensor interface {
	ReportData() string
}

// TemperatureSensor struct implementing the Sensor interface
type TemperatureSensor struct {
	temperature float64
}

// NewTemperatureSensor creates a new TemperatureSensor
func NewTemperatureSensor(temp float64) Sensor {
	return &TemperatureSensor{
		temperature: temp,
	}
}

// ReportData returns the temperature data
func (t *TemperatureSensor) ReportData() string {
	return fmt.Sprintf("Temperature: %.2f°C", t.temperature)
}

// HumiditySensor struct implementing the Sensor interface
type HumiditySensor struct {
	humidity float64
}

// NewHumiditySensor creates a new HumiditySensor
func NewHumiditySensor(humidity float64) Sensor {
	return &HumiditySensor{
		humidity: humidity,
	}
}

// ReportData returns the humidity data
func (h *HumiditySensor) ReportData() string {
	return fmt.Sprintf("Humidity: %.2f%%", h.humidity)
}

//SessionManager

// SensorManager interface with methods for managing sensors
type SensorManager interface {
	AddSensor(sensor Sensor)
	ReportAllData()
}

// BasicSensorManager struct managing a collection of sensors
type BasicSensorManager struct {
	sensors []Sensor
}

// NewBasicSensorManager creates a new BasicSensorManager
func NewBasicSensorManager() *BasicSensorManager {
	return &BasicSensorManager{
		sensors: []Sensor{},
	}
}

// AddSensor adds a new sensor to the manager
func (sm *BasicSensorManager) AddSensor(sensor Sensor) {
	sm.sensors = append(sm.sensors, sensor)
}

// ReportAllData reports data from all sensors managed by the manager
func (sm *BasicSensorManager) ReportAllData() {
	for _, sensor := range sm.sensors {
		fmt.Println(sensor.ReportData())
	}
}

// Controller
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

func main() {

	tempSensor := NewTemperatureSensor(23.5)
	humiditySensor := NewHumiditySensor(45.7)

	manager := NewBasicSensorManager()

	controller := NewSensorController(manager)

	controller.AddSensor(tempSensor)
	controller.AddSensor(humiditySensor)

	controller.ReportAllData()

}
