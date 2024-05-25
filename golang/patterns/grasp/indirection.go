package main

import "fmt"

type Service struct{}

func (s *Service) PerformOperation() {
	fmt.Println("Operation performed")
}

type Facade struct {
	service *Service
}

func NewFacade() *Facade {
	return &Facade{service: &Service{}}
}

func (f *Facade) PerformOperation() {
	f.service.PerformOperation()
}


func main() {
	facade := NewFacade()
	facade.PerformOperation()
}
