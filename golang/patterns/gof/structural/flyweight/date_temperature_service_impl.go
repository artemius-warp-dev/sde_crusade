package main

type DateTemperatureServiceImpl struct {
	flyweight *DateTemperatureFlyweight
}

func NewDateTemperatureServiceImpl(flyweight *DateTemperatureFlyweight) *DateTemperatureServiceImpl {

	return &DateTemperatureServiceImpl{flyweight: flyweight}

}

// GetDate returns the date from the flyweight.
func (d *DateTemperatureServiceImpl) GetDate() string {
	return d.flyweight.Date
}

// GetTemperature returns the temperature from the flyweight.
func (d *DateTemperatureServiceImpl) GetTemperature() float64 {
	return d.flyweight.Temperature
}
