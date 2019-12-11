package domain

import "fmt"

type PizzaBase1 interface {
	Prepare()
}

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type PizzaBase2 struct {
}

func (pizza *PizzaBase2) Bake() {
	fmt.Println("Let's bake")
}

func (pizza *PizzaBase2) Cut() {
	fmt.Println("Let's cut")
}

func (pizza *PizzaBase2) Box() {
	fmt.Println("Let's box")
}

type CheesePizza struct {
	PizzaBase2
}

func (pizza *CheesePizza) Prepare() {
	fmt.Println("Prepare cheese")
}

type GreekPizza struct {
	PizzaBase2
}

func (pizza *GreekPizza) Prepare() {
	fmt.Println("Prepare yogurt")
}
