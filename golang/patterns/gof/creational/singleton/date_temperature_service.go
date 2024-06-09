package main

import (
	"fmt"
	"sync"
)

// DateTemperatureService struct with DateSensor and TemperatureSensor
type DateTemperatureService struct {
	date        *DateSensor
	temperature *TemperatureSensor
}

// Singleton instance and sync.Once for thread-safe initialization
var instance *DateTemperatureService
var once sync.Once

// GetInstance returns the singleton instance of DateTemperatureService
func GetInstance() *DateTemperatureService {
	once.Do(func() {
		instance = &DateTemperatureService{
			date:        &DateSensor{date: "2024-06-08"},
			temperature: &TemperatureSensor{temperature: 36.6},
		}
	})
	return instance
}

// ReportData reports data from both DateSensor and TemperatureSensor
func (s *DateTemperatureService) ReportData() {
	fmt.Println(s.date.ReportData())
	fmt.Println(s.temperature.ReportData())
}
