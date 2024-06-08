package main

import "fmt"

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
