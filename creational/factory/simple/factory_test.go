package simple

import (
	"fmt"
	"testing"
)

var testList = []string{"cheese", "greek", "unknown"}

func TestSimpleFactory(t *testing.T) {
	order := SimpleOrder{}

	for _, test := range testList {
		fmt.Println("======================")
		fmt.Printf("Creating %s pizza \n", test)
		order.OrderPizza(test)
	}
}

func TestSimpleFactoryNoStruct(t *testing.T) {
	for _, test := range testList {
		fmt.Println("======================")
		fmt.Printf("Creating %s pizza \n", test)
		OrderPizzaNoStruct(test)
	}
}
