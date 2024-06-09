package main

// AbstractFactory interface
type AbstractFactory interface {
	CreateDateSensor() DateSensor
	CreateTemperatureSensor() TemperatureSensor
	CreateDateTemperatureService() *DateTemperatureService
}

// BasicSensorFactory is a concrete factory that produces basic sensors
type BasicSensorFactory struct{}

func (f *BasicSensorFactory) CreateDateSensor() DateSensor {
	return &BasicDateSensor{date: "2024-06-08"}
}

func (f *BasicSensorFactory) CreateTemperatureSensor() TemperatureSensor {
	return &BasicTemperatureSensor{temperature: 36.6}
}

func (f *BasicSensorFactory) CreateDateTemperatureService() *DateTemperatureService {
	return &DateTemperatureService{
		dateSensor:        f.CreateDateSensor(),
		temperatureSensor: f.CreateTemperatureSensor(),
	}
}
