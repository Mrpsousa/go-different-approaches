package main

import "fmt"

// Component Interface
type Coffee interface {
	Cost() float64
	Description() string
}

// Concrete Component that implement coffe interface
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// Decorator 1: Milk : Milk has coffee field that is Coffee interface type, then when we going to use Milk, we need to initialize this field with
// a struct that implements the Coffee interface then add new behavior if we want.
type Milk struct {
	coffee Coffee
}

//
func (m *Milk) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

func (m *Milk) Description() string {
	return m.coffee.Description() + ", Milk"
}

// Decorator 2: Caramel
type Caramel struct {
	coffee Coffee
}

func (c *Caramel) Cost() float64 {
	return c.coffee.Cost() + 1.0
}

func (c *Caramel) Description() string {
	return c.coffee.Description() + ", Caramel"
}

func main() {
	// Create a simple coffee
	coffee := &SimpleCoffee{}

	// Decorate it with Milk
	coffeeWithMilk := &Milk{coffee: coffee}
	fmt.Println("Coffee with Milk:", coffeeWithMilk.Description(), "- Cost:", coffeeWithMilk.Cost())

	// // Decorate it with Caramel
	coffeeWithCaramel := &Caramel{coffee: coffee}
	coffeeWithCaramel.coffee.Description()
	fmt.Println("Coffee with Caramel:", coffeeWithCaramel.Description(), "- Cost:", coffeeWithCaramel.Cost())

	// Decorate it with Milk and Caramel
	coffeeWithMilkAndCaramel := &Milk{coffee: &Caramel{coffee: coffee}}
	fmt.Println("Coffee with Milk and Caramel:", coffeeWithMilkAndCaramel.Description(), "- Cost:", coffeeWithMilkAndCaramel.Cost())
}
