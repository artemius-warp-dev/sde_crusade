package main

func main() {
	// Create a date temperature service
	service := &SimpleDateTemperatureService{}

	// Create commands
	getDateCommand := &GetDateCommand{service: service}
	getTemperatureCommand := &GetTemperatureCommand{service: service}
	resetCommand := &ResetCommand{service: service}

	// Create an invoker and add commands to it
	invoker := &Invoker{}
	invoker.AddCommand(getDateCommand)
	invoker.AddCommand(getTemperatureCommand)
	invoker.AddCommand(resetCommand)

	// Execute commands
	invoker.ExecuteCommands()

}
