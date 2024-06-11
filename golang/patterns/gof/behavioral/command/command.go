package main

import "fmt"

type Command interface {
	Execute()
}

type GetDateCommand struct {
	service DateTemperatureService
}

// Execute prints the date from the service.
func (c *GetDateCommand) Execute() {
	date := c.service.GetDate()
	fmt.Println("Date:", date)
}

// GetTemperatureCommand is a command to get the temperature from the service.
type GetTemperatureCommand struct {
	service DateTemperatureService
}

// Execute prints the temperature from the service.
func (c *GetTemperatureCommand) Execute() {
	temperature := c.service.GetTemperature()
	fmt.Println("Temperature:", temperature)
}

// ResetCommand is a command to reset the service.
type ResetCommand struct {
	service DateTemperatureService
}

// Execute resets the service.
func (c *ResetCommand) Execute() {
	c.service.Reset()
}
