package simple

import (
	"github.com/go_design_pattern/creational/factory/domain"
	"sync"
)
var simpleFactoryInstance simpleFactory
var once sync.Once

func GenSimpleFactory() simpleFactory {
	once.Do(func() {
		simpleFactoryInstance = simpleFactory{}
	})
	return simpleFactoryInstance
}

type simpleFactory struct {
}

func (factory *simpleFactory) GeneratePizza(pizzaType string) domain.Pizza {
	var pizza domain.Pizza
	switch pizzaType {
	case "cheese":
		pizza = &domain.CheesePizza{}
	case "greek":
		pizza = &domain.GreekPizza{}
	default:
		return nil
	}
	return pizza
}

// alternative without struct
func GeneratePizzaNoStruct(pizzaType string) domain.Pizza {
	var pizza domain.Pizza
	switch pizzaType {
	case "cheese":
		pizza = &domain.CheesePizza{}
	case "greek":
		pizza = &domain.GreekPizza{}
	default:
		return nil
	}
	return pizza
}
