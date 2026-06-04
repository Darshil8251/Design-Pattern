package main

import (
	"fmt"
)

type Pizza interface {
	GetDescription() string
	GetCost() float64
}

type plainPizza struct{}

func (p *plainPizza) GetDescription() string {
	return "Plain Pizza with tomato sauce and cheese"
}

func (p *plainPizza) GetCost() float64 {
	return 8.00
}

type margaritaPizza struct {
	pizza Pizza
}

func newMargaritaPizza(pizza Pizza) *margaritaPizza {
	return &margaritaPizza{
		pizza: pizza,
	}
}

func (p *margaritaPizza) GetDescription() string {
	return p.pizza.GetDescription() + ", margarita toppings (basil, mozzarella)"
}

func (p *margaritaPizza) GetCost() float64 {
	return p.pizza.GetCost() + 2.50
}

type farmvillaPizza struct {
	Pizza
}

func newFarmvillPizza(pizza Pizza) *farmvillaPizza {
	return &farmvillaPizza{
		Pizza: pizza,
	}
}

func (m *farmvillaPizza) GetDescription() string {
	return m.Pizza.GetDescription() + ", farm villa toppings (bell peppers, olives, onions)"
}

func (m *farmvillaPizza) GetCost() float64 {
	return m.Pizza.GetCost() + 1.80
}

// Helper function to print pizza details
func printPizza(name string, pizza Pizza) {
	fmt.Printf("%s:\n", name)
	fmt.Printf("Description: %s\n", pizza.GetDescription())
	fmt.Printf("Cost: $%.2f\n\n", pizza.GetCost())
}

// Pizza ordering system demonstration
func main() {
	plain := &plainPizza{}
	printPizza("Plain Pizza", plain)

	margarita := newMargaritaPizza(plain)
	printPizza("Margarita Pizza", margarita)

	farmVilla := newFarmvillPizza(plain)
	printPizza("Farm Villa Pizza", farmVilla)
	fmt.Println("Enjoy your pizzas!")
}
