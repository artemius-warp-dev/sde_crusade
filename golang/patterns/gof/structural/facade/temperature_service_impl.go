package main

type TemperatureServiceImpl struct{}

func (tsi TemperatureServiceImpl) GetTemperature() float64 {
	return 25.5
}
