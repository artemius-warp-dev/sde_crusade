package main

type DateTemperatureFacade struct {
	dateService        DateService
	temperatureService TemperatureService
}

type DateService interface {
	GetDate() string
}

type TemperatureService interface {
	GetTemperature() float64
}

func NewDateTemperatureFacade(dateService DateService, temperatureService TemperatureService) *DateTemperatureFacade {

	return &DateTemperatureFacade{
		dateService:        dateService,
		temperatureService: temperatureService,
	}
}

func (dtf *DateTemperatureFacade) GetDate() string {
	return dtf.dateService.GetDate()
}

func (dtf *DateTemperatureFacade) GetTemperature() float64 {
	return dtf.temperatureService.GetTemperature()
}
