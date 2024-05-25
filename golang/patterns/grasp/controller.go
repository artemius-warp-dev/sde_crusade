package main

import "fmt"

type Controller struct{}

func (c *Controller) HandleEvent(event string) {
	fmt.Println("Event handled", event)
}

func main() {
	controller := Controller{}
	controller.HandleEvent("UserClick")
}
