package main

import "fmt"

// LoggingHandler logs the service data.
type LoggingHandler struct {
	BaseHandler
}

// Handle logs the date and temperature.
func (h *LoggingHandler) Handle(service DateTemperatureService) {
	date := service.GetDate()
	temperature := service.GetTemperature()

	fmt.Printf("Logging data - Date: %s, Temperature: %.1f\n", date, temperature)
	h.BaseHandler.Handle(service)
}
