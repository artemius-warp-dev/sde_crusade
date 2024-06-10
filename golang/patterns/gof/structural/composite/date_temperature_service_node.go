package main

type DateTemperatureServiceNode struct {
	date        string
	temperature float64
}

func NewDateTemperatureServiceNode(date string, temperature float64) *DateTemperatureServiceNode {

	return &DateTemperatureServiceNode{date: date, temperature: temperature}

}

func (s *DateTemperatureServiceNode) GetDate() string {

	return s.date

}

func (s *DateTemperatureServiceNode) GetTemperature() float64 {
	return s.temperature
}
