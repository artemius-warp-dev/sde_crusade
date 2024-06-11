package main

func main() {
	// Create a date temperature service
	service := &SimpleDateTemperatureService{}

	// Create handlers
	validationHandler := &ValidationHandler{}
	processingHandler := &ProcessingHandler{}
	loggingHandler := &LoggingHandler{}

	// Set up the chain of responsibility
	validationHandler.SetNext(processingHandler).SetNext(loggingHandler)

	// Start the chain
	validationHandler.Handle(service)

}
