package main

import (
	"fmt"
	"sync"
)

// DateTemperatureFlyweight represents shared date and temperature data.
type DateTemperatureFlyweight struct {
	Date        string
	Temperature float64
}

type DateTemperatureFactory struct {
	flyweights map[string]*DateTemperatureFlyweight
	mutex      sync.Mutex
}

func NewDateTemperatureFactory() *DateTemperatureFactory {

	return &DateTemperatureFactory{
		flyweights: make(map[string]*DateTemperatureFlyweight),
	}
}

func (f *DateTemperatureFactory) GetFlyweight(date string, temperature float64) *DateTemperatureFlyweight {

	f.mutex.Lock()
	defer f.mutex.Unlock()

	key := fmt.Sprintf("%s:%f", date, temperature)

	if flyweight, exists := f.flyweights[key]; exists {
		return flyweight
	}

	flyweight := &DateTemperatureFlyweight{
		Date:        date,
		Temperature: temperature,
	}

	f.flyweights[key] = flyweight
	return flyweight

}
