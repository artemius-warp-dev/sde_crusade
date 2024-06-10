package main

type DateTemperatureService interface {
	GetDate() string
	GetTemperature() float64
}
