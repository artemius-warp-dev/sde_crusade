package main

func main() {

	// Create a date temperature service
	service := &SimpleDateTemperatureService{}

	// Create a mediator
	mediator := NewConcreteMediator(service)

	// Create components
	dateComponent := &DateComponent{}
	temperatureComponent := &TemperatureComponent{}
	resetComponent := &ResetComponent{}
	loggingComponent := &LoggingComponent{}

	// Register components with the mediator
	mediator.RegisterComponent("DateComponent", dateComponent)
	mediator.RegisterComponent("TemperatureComponent", temperatureComponent)
	mediator.RegisterComponent("ResetComponent", resetComponent)
	mediator.RegisterComponent("LoggingComponent", loggingComponent)

	// Use components to interact via the mediator
	dateComponent.GetDate()
	temperatureComponent.GetTemperature()
	loggingComponent.Log()
	resetComponent.Reset()
	loggingComponent.Log()

}
