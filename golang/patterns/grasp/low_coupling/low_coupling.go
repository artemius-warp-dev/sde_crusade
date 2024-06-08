package main

import "fmt"

type Goods interface {
	GetName() string
	GetPrice() float64
}

type ConcreteGoods struct {
	name  string
	price float64
}

func NewConcreteGoods(name string, price float64) Goods {
	return &ConcreteGoods{
		name:  name,
		price: price,
	}
}

func (g *ConcreteGoods) GetName() string {
	return g.name
}

func (g *ConcreteGoods) GetPrice() float64 {
	return g.price

}

type OrderItem struct {
	goods     Goods
	quantity  int
	totalCost float64
}

// NewOrderItem is a factory method for creating OrderItem
func NewOrderItem(goods Goods, quantity int) OrderItem {
	return OrderItem{
		goods:     goods,
		quantity:  quantity,
		totalCost: goods.GetPrice() * float64(quantity),
	}
}

func (item *OrderItem) GetTotalCost() float64 {
	return item.totalCost
}

func (item *OrderItem) GetGoods() Goods {
	return item.goods
}

type Order struct {
	items []OrderItem
}

func NewOrder() Order {
	return Order{
		items: []OrderItem{},
	}
}

func (o *Order) AddItem(item OrderItem) {
	o.items = append(o.items, item)
}

func (o *Order) TotalCost() float64 {
	total := 0.0
	for _, item := range o.items {
		total += item.GetTotalCost()
	}

	return total
}

func main() {
	// Creating goods
	goods1 := NewConcreteGoods("Apple", 0.50)
	goods2 := NewConcreteGoods("Banana", 0.30)

	// Creating an order
	order := NewOrder()

	// Adding items to the order
	order.AddItem(NewOrderItem(goods1, 10)) // 10 apples
	order.AddItem(NewOrderItem(goods2, 20)) // 20 bananas

	// Calculating total cost
	fmt.Printf("Total cost of the order: $%.2f\n", order.TotalCost())
}
