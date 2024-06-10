package main

import "fmt"

func main() {

	service1 := NewDateTemperatureServiceNode("2024-06-09", 25.5)
	service2 := NewDateTemperatureServiceNode("2024-06-09", 26.0)
	service3 := NewDateTemperatureServiceNode("2024-06-09", 24.5)

	composite1 := NewCompositeDateTemperatureService()
	composite2 := NewCompositeDateTemperatureService()

	mainComposite := NewCompositeDateTemperatureService()

	// Build the composite structure (tree)
	composite1.Add(service1)
	composite1.Add(service2)
	composite2.Add(service3)

	mainComposite.Add(composite1)
	mainComposite.Add(composite2)

	// Use the main composite service
	fmt.Println("Main Composite Date:", mainComposite.GetDate())
	fmt.Println("Main Composite Average Temperature:", mainComposite.GetTemperature())

}
