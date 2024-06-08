package main

import (
	"fmt"
	"time"
)

type Date interface {
	CurrentDate() string
}

type Temperature interface {
	CurrentTemperature() float64
}

type DateTempDevice struct {
	date Date
	temp Temperature
}

type ConcreteDate struct{}

func (d ConcreteDate) CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

type ConcreteTemperature struct {
	temperature float64
}

func newConcreteTemperature(temp float64) Temperature {
	return &ConcreteTemperature{temperature: temp}
}

func (t ConcreteTemperature) CurrentTemperature() float64 {
	return t.temperature
}

func NewDateTempDevice(date Date, temp Temperature) DateTempDevice {
	return DateTempDevice{
		date: date,
		temp: temp,
	}
}

func (d DateTempDevice) GetDateAndTemperature() string {
	return fmt.Sprintf("Date: %s, Temperature: %.2f C", d.date.CurrentDate(), d.temp.CurrentTemperature())
}

func main() {
	date := ConcreteDate{}
	temp := newConcreteTemperature(23.5)

	device := NewDateTempDevice(date, temp)

	fmt.Println(device.GetDateAndTemperature())
}
