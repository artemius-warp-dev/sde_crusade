package main

import "fmt"

func main() {

	dateService := DateServiceImpl{}
	temperatureService := TemperatureServiceImpl{}

	facade := NewDateTemperatureFacade(dateService, temperatureService)

	fmt.Println("Date:", facade.GetDate())
	fmt.Println("Temperature:", facade.GetTemperature())

}
