package main

import "fmt"

type Sensor interface {
	ReportData() string
}

type TemperatureSensor struct {
	temperature float64
}

// NewTemperatureSensor creates a new TemperatureSensor
func NewTemperatureSensor(temp float64) Sensor {
	return &TemperatureSensor{
		temperature: temp,
	}
}

// ReportData returns the temperature data
func (t *TemperatureSensor) ReportData() string {
	return fmt.Sprintf("Temperature: %.2f°C", t.temperature)
}

// HumiditySensor struct implementing the Sensor interface
type HumiditySensor struct {
	humidity float64
}

// NewHumiditySensor creates a new HumiditySensor
func NewHumiditySensor(humidity float64) Sensor {
	return &HumiditySensor{
		humidity: humidity,
	}
}

// ReportData returns the humidity data
func (h *HumiditySensor) ReportData() string {
	return fmt.Sprintf("Humidity: %.2f%%", h.humidity)
}

func main() {
	tempSensor := NewTemperatureSensor(23.5)
	humiditySensor := NewHumiditySensor(45.7)

	sensors := []Sensor{tempSensor, humiditySensor}

	for _, sensor := range sensors {
		fmt.Println(sensor.ReportData())
	}
}
