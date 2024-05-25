package main

import "fmt"

type Product struct {
	Name string
	Price float64
}

type ProductService struct{}

func (ps *ProductService) CalculateTotalPrice(products []*Product) float64 {
	total := 0.0
	for _, p := range products {
		total += p.Price
	}

	return total
}

func main() {
	products := []*Product{
		{Name: "Laptop", Price: 1000.00},
		{Name: "Phone", Price: 5000},
	}
	productService := ProductService{}
	totalPrice := productService.CalculateTotalPrice(products)
	fmt.Println("Total Price:", totalPrice)
}
