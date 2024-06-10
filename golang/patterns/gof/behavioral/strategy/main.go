package main

import "fmt"

func main() {

	service := &SimpleDateTemperatureService{}

	simpleStrategy := &SimpleFormatStrategy{}

	detailedStrategy := &DetailedFormatStrategy{}

	context := NewContext(service, simpleStrategy)
	fmt.Println("Simple Strategy Output:", context.GetFormattedData())

	context.SetStrategy(detailedStrategy)
	fmt.Println("Detailed Strategy Output:", context.GetFormattedData())

}
