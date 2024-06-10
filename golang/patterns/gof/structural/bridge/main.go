package main

import "fmt"

func main() {

	simpleImplementor := &SimpleDateTemperatureImplementor{}
	apiImplementor := &APIDateTemperatureImplementor{}
	databaseImplementor := &DatabaseDateTemperatureImplementor{}

	// Initialize the bridge services with different implementors
	simpleService := NewBridgeDateTemperatureService(simpleImplementor)
	apiService := NewBridgeDateTemperatureService(apiImplementor)
	databaseService := NewBridgeDateTemperatureService(databaseImplementor)

	// Use the services
	fmt.Println("Simple Service Date:", simpleService.GetDate())
	fmt.Println("Simple Service Temperature:", simpleService.GetTemperature())

	fmt.Println("API Service Date:", apiService.GetDate())
	fmt.Println("API Service Temperature:", apiService.GetTemperature())

	fmt.Println("Database Service Date:", databaseService.GetDate())
	fmt.Println("Database Service Temperature:", databaseService.GetTemperature())

}
