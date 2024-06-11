package main

import "fmt"

type Mediator interface {
	Notify(sender Component, event string)
}

type ConcreteMediator struct {
	service    DateTemperatureService
	components map[string]Component
}

func NewConcreteMediator(service DateTemperatureService) *ConcreteMediator {
	return &ConcreteMediator{
		service:    service,
		components: make(map[string]Component),
	}
}

func (m *ConcreteMediator) RegisterComponent(name string, component Component) {
	component.SetMediator(m)
	m.components[name] = component
}

func (m *ConcreteMediator) Notify(sender Component, event string) {
	switch event {
	case "GetDate":
		fmt.Println("Mediator: Getting date")
		fmt.Println("Date:", m.service.GetDate())
	case "GetTemperature":
		fmt.Println("Mediator: Getting temperature")
		fmt.Println("Temperature:", m.service.GetTemperature())
	case "Reset":
		fmt.Println("Mediator: Resetting service")
		m.service.Reset()
	case "Log":
		fmt.Println("Mediator: Logging date")
		date := m.service.GetDate()
		temperature := m.service.GetTemperature()
		fmt.Println("Logging Date:", date, "Temperature:", temperature)
	}
}
