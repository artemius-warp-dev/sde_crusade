package main

type Invoker struct {
	commands []Command
}

func (i *Invoker) AddCommand(command Command) {

	i.commands = append(i.commands, command)

}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}
