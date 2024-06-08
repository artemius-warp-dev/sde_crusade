package main

import (
	"fmt"
)

// Goods struct representing the goods that can be ordered
type Goods struct {
	Name  string
	Price float64
}

// NewGoods is a factory method for creating Goods
func NewGoods(name string, price float64) Goods {
	return Goods{
		Name:  name,
		Price: price,
	}
}

// OrderItem struct representing an item in an order
type OrderItem struct {
	Goods     Goods
	Quantity  int
	TotalCost float64
}

// NewOrderItem is a factory method for creating OrderItem
func NewOrderItem(goods Goods, quantity int) OrderItem {
	return OrderItem{
		Goods:     goods,
		Quantity:  quantity,
		TotalCost: goods.Price * float64(quantity),
	}
}

// Order struct representing the order itself
type Order struct {
	Items []OrderItem
}

// NewOrder is a factory method for creating an Order
func NewOrder() Order {
	return Order{
		Items: []OrderItem{},
	}
}

// AddItem adds a new OrderItem to the Order
func (o *Order) AddItem(goods Goods, quantity int) {
	item := NewOrderItem(goods, quantity)
	o.Items = append(o.Items, item)
}

// TotalCost calculates the total cost of the order
func (o *Order) TotalCost() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.TotalCost
	}
	return total
}

func main() {
	// Creating goods
	goods1 := NewGoods("Apple", 0.50)
	goods2 := NewGoods("Banana", 0.30)

	// Creating an order
	order := NewOrder()

	// Adding items to the order
	order.AddItem(goods1, 10) // 10 apples
	order.AddItem(goods2, 20) // 20 bananas

	// Calculating total cost
	fmt.Printf("Total cost of the order: $%.2f\n", order.TotalCost())
}
