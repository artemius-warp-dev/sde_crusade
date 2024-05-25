package main

import "fmt"

type ServiceInterface interface {
	PerformAction()
}

type ServiceA struct{}

func (s *ServiceA) PerformAction() {
	fmt.Println("ServiceA action performed")
}


type ServiceB struct{}

func (s *ServiceB) PerformAction() {
	fmt.Println("ServiceB action performed")
}


type Client struct {
	service ServiceInterface
}


func NewClient(service ServiceInterface) *Client {
	return &Client{service: service}
}

func (c *Client) DoAction() {
	c.service.PerformAction()
}

func main() {
	serviceA := &ServiceA{}
	serviceB := &ServiceB{}

	client1 := NewClient(serviceA)
	client1.DoAction()

	client2 := NewClient(serviceB)
	client2.DoAction()
}
