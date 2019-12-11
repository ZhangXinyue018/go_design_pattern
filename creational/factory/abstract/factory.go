package abstract

import (
	"fmt"
	"github.com/go_design_pattern/creational/factory/domain"
)

type HardCheesePizza struct {
	domain.PizzaBase2
}

func (pizza *HardCheesePizza) Prepare() {
	fmt.Println("Prepare hard cheese")
}

type SoftCheesePizza struct {
	domain.PizzaBase2
}

func (pizza *SoftCheesePizza) Prepare() {
	fmt.Println("Prepare soft cheese")
}

type NorthGreekPizza struct {
	domain.PizzaBase2
}

func (pizza *NorthGreekPizza) Prepare() {
	fmt.Println("Prepare north yogurt")
}

type SouthGreekPizza struct {
	domain.PizzaBase2
}

func (pizza *SouthGreekPizza) Prepare() {
	fmt.Println("Prepare south yogurt")
}

// abstract layer
type Factory interface {
	GeneratePizza(pizzaType string) domain.Pizza
}

type CheeseFactory struct {
}

func (factory *CheeseFactory) GeneratePizza(pizzaType string) domain.Pizza {
	switch pizzaType {
	case "soft":
		return &SoftCheesePizza{}
	case "hard":
		return &HardCheesePizza{}
	default:
		return &domain.CheesePizza{}
	}
}

type GreekFactory struct {
}

func (factory *GreekFactory) GeneratePizza(pizzaType string) domain.Pizza {
	switch pizzaType {
	case "north":
		return &NorthGreekPizza{}
	case "south":
		return &SouthGreekPizza{}
	default:
		return &domain.GreekPizza{}
	}
}
