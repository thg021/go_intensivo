package entity

import (
	"errors"
	"fmt"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// Regra de negocio
func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}

	return order, nil
}

// "metodo" da struct
func (o *Order) Validate() error {

	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0 {
		fmt.Println("Aqui")
		return errors.New("invalid price")
	}

	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	err := o.Validate()
	if err != nil {
		return err
	}

	return nil
}
