package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(service DateTemperatureService)
}

type GetDateExpression struct{}

func (e *GetDateExpression) Interpret(service DateTemperatureService) {

	date := service.GetDate()
	fmt.Println("Date:", date)

}

type GetTemperatureExpression struct{}

func (e *GetTemperatureExpression) Interpret(service DateTemperatureService) {
	temperature := service.GetTemperature()
	fmt.Println("Temperature:", temperature)
}

type ResetExpression struct{}

func (e *ResetExpression) Interpret(service DateTemperatureService) {
	service.Reset()
}

type Interpreter struct {
	service DateTemperatureService
}

func NewInterpreter(service DateTemperatureService) *Interpreter {

	return &Interpreter{service: service}
}

func (i *Interpreter) Interpret(input string) {
	commands := strings.Fields(input)
	for _, command := range commands {
		var expression Expression
		switch command {
		case "date":
			expression = &GetDateExpression{}
		case "temperature":
			expression = &GetTemperatureExpression{}
		case "reset":
			expression = &ResetExpression{}
		default:
			fmt.Println("Unknown command:", command)
			continue
		}
		expression.Interpret(i.service)

	}
}
