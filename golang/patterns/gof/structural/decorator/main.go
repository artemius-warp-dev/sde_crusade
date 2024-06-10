package main

import "fmt"

func main() {

	realService := &OldDateTemperatureService{}

	loggingDecorator := NewLoggingDecorator(realService)
	cachingDecorator := NewCachingDecorator(loggingDecorator)

	fmt.Println("Date:", cachingDecorator.GetDate())
	fmt.Println("Temperature:", cachingDecorator.GetTemperature())

	fmt.Println("Date (cached):", cachingDecorator.GetDate())
	fmt.Println("Temperature (cached):", cachingDecorator.GetTemperature())

}
