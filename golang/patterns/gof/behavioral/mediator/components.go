package main

type Component interface {
	SetMediator(mediator Mediator)
}

type BaseComponent struct {
	mediator Mediator
}

func (c *BaseComponent) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

type DateComponent struct {
	BaseComponent
}

func (c *DateComponent) GetDate() {
	c.mediator.Notify(c, "GetDate")
}

type TemperatureComponent struct {
	BaseComponent
}

func (c *TemperatureComponent) GetTemperature() {
	c.mediator.Notify(c, "GetTemperature")
}

type ResetComponent struct {
	BaseComponent
}

func (c *ResetComponent) Reset() {
	c.mediator.Notify(c, "Reset")
}

type LoggingComponent struct {
	BaseComponent
}

func (c *LoggingComponent) Log() {
	c.mediator.Notify(c, "Log")
}
