package main

type ServiceBuilder interface {
	SetDate(date string) ServiceBuilder
	SetTemperature(temp float64) ServiceBuilder
	Build() *DateTemperatureService
}

type DateTemperatureServiceBuilder struct {
	date        string
	temperature float64
}

func (b *DateTemperatureServiceBuilder) SetDate(date string) ServiceBuilder {
	b.date = date
	return b
}

func (b *DateTemperatureServiceBuilder) SetTemperature(temp float64) ServiceBuilder {
	b.temperature = temp
	return b
}

func (b *DateTemperatureServiceBuilder) Build() *DateTemperatureService {
	return &DateTemperatureService{
		date:        &DateSensor{date: b.date},
		temperature: &TemperatureSensor{temperature: b.temperature},
	}
}
