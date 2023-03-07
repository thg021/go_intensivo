package main

import "fmt"

// Seria a Class
type Order struct {
	ID       string
	Price    float64
	Quantity int
}

// Seria o Metodo
func (o *Order) setPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do setPrice : ", o.Price)
}

func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

// Seria o constructor
func NewOrder() *Order {
	return &Order{}
}

func main() {
	order := NewOrder()

	order.ID = "123"
	order.Price = 10.0
	order.Quantity = 5

	order.setPrice(20.0)
	fmt.Println("Preco total: ", order.getTotal())
}
