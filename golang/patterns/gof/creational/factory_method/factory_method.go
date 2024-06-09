package main

type ServiceFactory interface {
	CreateService() *DateTemperatureService
}

type DateTemperatureServiceFactory struct{}

func (f *DateTemperatureServiceFactory) CreateService() *DateTemperatureService {
	return &DateTemperatureService{
		date:        &DateSensor{date: "2024-06-08"},
		temperature: &TemperatureSensor{temperature: 36.6},
	}
}
