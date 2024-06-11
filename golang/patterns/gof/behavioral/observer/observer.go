package main

import "fmt"

type Observer interface {
	Update(service DateTemperatureService)
}

type DateObserver struct{}

func (o *DateObserver) Update(service DateTemperatureService) {
	fmt.Println("DateObserver: Date updated to", service.GetDate())
}

type TemperatureObserver struct{}

func (o *TemperatureObserver) Update(service DateTemperatureService) {
	fmt.Println("TemperatureObserver: Temperature updated to", service.GetTemperature())
}
