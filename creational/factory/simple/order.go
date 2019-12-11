package simple

import "fmt"

type SimpleOrder struct {
}

func (order *SimpleOrder) OrderPizza(pizzaType string) {
	factory := GenSimpleFactory()
	pizza := factory.GeneratePizza(pizzaType)
	if pizza != nil {
		pizza.Prepare()
		pizza.Bake()
		pizza.Cut()
		pizza.Box()
	} else {
		fmt.Println("Invalid pizza")
	}
}

// alternative without struct
func OrderPizzaNoStruct(pizzaType string) {
	pizza := GeneratePizzaNoStruct(pizzaType)
	if pizza != nil {
		pizza.Prepare()
		pizza.Bake()
		pizza.Cut()
		pizza.Box()
	} else {
		fmt.Println("Invalid pizza")
	}
}
