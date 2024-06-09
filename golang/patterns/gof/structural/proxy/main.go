package main

import "fmt"

func main() {

	realService := OldDateTemperatureService{}

	loggingProxy := &LoggingProxy{RealService: realService}

	// Initialize the auth proxy with the logging proxy
	authProxy := AuthProxy{
		RealService: loggingProxy,
		Username:    "admin",
		Password:    "password",
	}

	fmt.Println("Date:", authProxy.DateSensor())
	fmt.Println("Temperature:", authProxy.TemperatureSensor())

}
