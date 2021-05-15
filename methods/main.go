package main

import (
	"errors"
	"fmt"
	"github.com/Rhymond/go-money"
	"log"
)

type Item struct {
	ID    string
	Name  string
	Price *money.Money
}

type Cart struct {
	ID       string
	IsLocked bool
	Items    []Item
}

func (c Cart) TotalPrice() (*money.Money, error) {
	totalPrice := money.New(0, "EUR")
	for _, item := range c.Items {
		var err error
		totalPrice, err = totalPrice.Add(item.Price)
		if err != nil {
			return nil, fmt.Errorf("impossible to compute total price: %w", err)
		}
	}
	return totalPrice, nil
}

func (c *Cart) Lock() error {
	if c.IsLocked {
		return errors.New("the cart is already locked, cannot be locked")
	}
	c.IsLocked = true
	return nil
}

func main() {
	items := []Item{
		{
			ID:    "458",
			Name:  "Book",
			Price: money.New(1000, "EUR"),
		},
		{
			ID:    "8888",
			Name:  "Book 2",
			Price: money.New(1200, "EUR"),
		},
		{
			ID:    "8888555",
			Name:  "Book 3",
			Price: money.New(1000, "EUR"),
		},
	}
	cart := Cart{
		ID:       "42",
		Items:    items,
		IsLocked: true,
	}
	total, err := cart.TotalPrice()
	if err != nil {
		log.Fatalf("Error while computing total price: %s", err)
	}

	fmt.Println(total.Display())
	err = cart.Lock()
	if err != nil {
		log.Fatalf("Error while locking the cart: %s", err)
	}
	fmt.Println("Cart is locked !")

}
