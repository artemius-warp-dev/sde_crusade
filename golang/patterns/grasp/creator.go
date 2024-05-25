package main

import "fmt"

type Product struct {
	Name string
	Price float64
}

type Factory struct{}

func (f *Fctory) CreateProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

func main() {
	factory := Factory{}
	product := factory.CreateProduct("Laptop", 1000.00)
	fmt.Println(product)
}
