package main

import "fmt"

func main() {
	// Initialize the flyweight factory
	factory := NewDateTemperatureFactory()

	// Get shared flyweight instances
	flyweight1 := factory.GetFlyweight("2024-06-09", 25.5)
	flyweight2 := factory.GetFlyweight("2024-06-09", 25.5)
	flyweight3 := factory.GetFlyweight("2024-06-10", 26.0)

	// Initialize the services using the flyweights
	service1 := NewDateTemperatureServiceImpl(flyweight1)
	service2 := NewDateTemperatureServiceImpl(flyweight2)
	service3 := NewDateTemperatureServiceImpl(flyweight3)

	// Use the services
	fmt.Println("Service1 - Date:", service1.GetDate(), "Temperature:", service1.GetTemperature())
	fmt.Println("Service2 - Date:", service2.GetDate(), "Temperature:", service2.GetTemperature())
	fmt.Println("Service3 - Date:", service3.GetDate(), "Temperature:", service3.GetTemperature())

	// Verify that the flyweights are shared
	fmt.Println("Flyweight1 and Flyweight2 are the same instance:", flyweight1 == flyweight2)
	fmt.Println("Flyweight1 and Flyweight3 are the same instance:", flyweight1 == flyweight3)

}
