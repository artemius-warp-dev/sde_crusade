package main

import "fmt"

// LoggingProxy represents a proxy for the DateTemperatureService with logging.
type LoggingProxy struct {
	RealService DateTemperatureService
}

// DateSensor logs the call and proxies it to the real service's DateSensor method.
func (lp *LoggingProxy) DateSensor() string {
	fmt.Println("LoggingProxy: Calling DateSensor")
	result := lp.RealService.DateSensor()
	fmt.Println("LoggingProxy: DateSensor returned", result)
	return result
}

// TemperatureSensor logs the call and proxies it to the real service's TemperatureSensor method.
func (lp *LoggingProxy) TemperatureSensor() float64 {
	fmt.Println("LoggingProxy: Calling TemperatureSensor")
	result := lp.RealService.TemperatureSensor()
	fmt.Println("LoggingProxy: TemperatureSensor returned", result)
	return result
}
