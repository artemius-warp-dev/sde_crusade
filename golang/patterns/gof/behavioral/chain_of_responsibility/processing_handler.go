package main

import "fmt"

type ProcessingHandler struct {
	BaseHandler
}

func (h *ProcessingHandler) Handle(service DateTemperatureService) {

	date := service.GetDate()
	temperature := service.GetTemperature()

	fmt.Printf("Processing data - Date: %s, Temperature: %.1f\n", date, temperature)
	h.BaseHandler.Handle(service)

}
