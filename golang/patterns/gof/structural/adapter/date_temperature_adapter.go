// date_temperature_adapter.go
package main

// DateTemperatureAdapter adapts the OldDateTemperatureService to the NewDateTemperatureService interface.
type DateTemperatureAdapter struct {
	OldService OldDateTemperatureService
}

// GetDate adapts the DateSensor method to GetDate.
func (dta DateTemperatureAdapter) GetDate() string {
	return dta.OldService.DateSensor()
}

// GetTemperature adapts the TemperatureSensor method to GetTemperature.
func (dta DateTemperatureAdapter) GetTemperature() float64 {
	return dta.OldService.TemperatureSensor()
}
