package main

import "fmt"

type Memento struct {
	date        string
	temperature float64
}

type DateTemperatureService struct {
	date        string
	temperature float64
}

// NewDateTemperatureService creates a new DateTemperatureService.
func NewDateTemperatureService(date string, temperature float64) *DateTemperatureService {
	return &DateTemperatureService{
		date:        date,
		temperature: temperature,
	}
}

// SetState sets the state of the DateTemperatureService.
func (s *DateTemperatureService) SetState(date string, temperature float64) {
	s.date = date
	s.temperature = temperature
}

// Save creates a memento containing the current state of the DateTemperatureService.
func (s *DateTemperatureService) Save() *Memento {
	return &Memento{
		date:        s.date,
		temperature: s.temperature,
	}
}

// Restore restores the DateTemperatureService's state from a memento.
func (s *DateTemperatureService) Restore(m *Memento) {
	s.date = m.date
	s.temperature = m.temperature
}

// Report prints the current state of the DateTemperatureService.
func (s *DateTemperatureService) Report() {
	fmt.Printf("Date: %s, Temperature: %.1f\n", s.date, s.temperature)
}
