package main

import "fmt"

func main() {
	// Initialize the old service
	oldService := OldDateTemperatureService{}

	// Initialize the adapter
	adapter := DateTemperatureAdapter{OldService: oldService}

	// Use the new service interface
	fmt.Println("Date:", adapter.GetDate())
	fmt.Println("Temperature:", adapter.GetTemperature())
}
