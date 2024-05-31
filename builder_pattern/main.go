package main

import "fmt"

// Obj to build
type Pizza struct {
	Size      string
	Crust     string
	Cheese    bool
	Pepperoni bool
	Mushrooms bool
	// other toppings...
}

// Creates builder for Pizza struct
// interface optional
type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	AddCheese() PizzaBuilder
	AddPepperoni() PizzaBuilder
	AddMushrooms() PizzaBuilder
	Build() Pizza
}

type ConcretePizzaBuilder struct {
	pizza Pizza
}

func (b *ConcretePizzaBuilder) SetSize(size string) PizzaBuilder {
	b.pizza.Size = size
	return b
}

func (b *ConcretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
	b.pizza.Crust = crust
	return b
}

func (b *ConcretePizzaBuilder) AddCheese() PizzaBuilder {
	b.pizza.Cheese = true
	return b
}

func (b *ConcretePizzaBuilder) AddPepperoni() PizzaBuilder {
	b.pizza.Pepperoni = true
	return b
}

func (b *ConcretePizzaBuilder) AddMushrooms() PizzaBuilder {
	b.pizza.Mushrooms = true
	return b
}

func (b *ConcretePizzaBuilder) Build() Pizza {
	return b.pizza
}

//main
func main() {
	builder := &ConcretePizzaBuilder{}

	// Custom Pizza
	cP := builder.SetSize("Large").AddPepperoni().AddMushrooms().SetCrust("chocolate").Build()
	fmt.Printf("Custom Pizza. Size:%s, Crust:%s, Chess:%v, Pepperoni:%v, Mushrooms:%v,  ", cP.Size, cP.Crust, cP.Cheese, cP.Pepperoni, cP.Mushrooms )

}
