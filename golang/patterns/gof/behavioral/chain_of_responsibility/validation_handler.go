package main

import "fmt"

type ValidationHandler struct {
	BaseHandler
}

func (h *ValidationHandler) Handle(service DateTemperatureService) {

	date := service.GetDate()
	temperature := service.GetTemperature()

	if date == "" || temperature == 0 {
		fmt.Println("Validation failed")
		return
	}

	fmt.Println("Validation passed.")
	h.BaseHandler.Handle(service)

}
