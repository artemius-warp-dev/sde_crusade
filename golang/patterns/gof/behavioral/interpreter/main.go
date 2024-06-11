package main

func main() {

	service := &SimpleDateTemperatureService{}

	interpreter := NewInterpreter(service)

	interpreter.Interpret("date temperature reset")

}
