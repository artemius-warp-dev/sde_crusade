package main

import "fmt"

// State interface represents the state of the service.
type State interface {
	GetDate(service *DateTemperatureServiceContext) string
	GetTemperature(service *DateTemperatureServiceContext) float64
	Reset(service *DateTemperatureServiceContext)
}

// NormalState represents the normal state of the service.
type NormalState struct {
	service     *DateTemperatureServiceContext
	date        string
	temperature float64
}

// NewNormalState creates a new NormalState.
func NewNormalState(service *DateTemperatureServiceContext, date string, temperature float64) *NormalState {
	return &NormalState{service: service, date: date, temperature: temperature}
}

// GetDate returns the current date in the normal state.
func (s *NormalState) GetDate(service *DateTemperatureServiceContext) string {
	return s.date
}

// GetTemperature returns the current temperature in the normal state.
func (s *NormalState) GetTemperature(service *DateTemperatureServiceContext) float64 {
	return s.temperature
}

// Reset resets the service to the initial state in the normal state.
func (s *NormalState) Reset(service *DateTemperatureServiceContext) {
	s.date = "2024-01-01"
	s.temperature = 0
	fmt.Println("Service has been reset to normal state.")
}

// MaintenanceState represents the maintenance state of the service.
type MaintenanceState struct {
	service *DateTemperatureServiceContext
}

// NewMaintenanceState creates a new MaintenanceState.
func NewMaintenanceState(service *DateTemperatureServiceContext) *MaintenanceState {
	return &MaintenanceState{service: service}
}

// GetDate returns a maintenance message for the date.
func (s *MaintenanceState) GetDate(service *DateTemperatureServiceContext) string {
	return "Service is under maintenance."
}

// GetTemperature returns a maintenance message for the temperature.
func (s *MaintenanceState) GetTemperature(service *DateTemperatureServiceContext) float64 {
	return -1
}

// Reset resets the service to the normal state.
func (s *MaintenanceState) Reset(service *DateTemperatureServiceContext) {
	fmt.Println("Switching to normal state.")
	s.service.SetState(NewNormalState(s.service, "2024-06-09", 25.5))
}
