package method

import (
	"fmt"
	"testing"
)

func TestCheeseFactoryMethod(t *testing.T) {
	fmt.Println("~~~~~~~~~~~~~~~Start cheese factory test~~~~~~~~~~~~~~~~")
	testList := []string{"soft", "hard", "others"}
	cheeseFactory := &CheeseFactory{}
	for _, test := range testList {
		fmt.Println("======================")
		fmt.Printf("Creating %s pizza \n", test)
		pizza := cheeseFactory.GeneratePizza(test)
		if pizza != nil {
			pizza.Prepare()
			pizza.Bake()
			pizza.Cut()
			pizza.Box()
		} else {
			fmt.Println("Invalid pizza")
		}
	}

}

func TestGreekFactoryMethod(t *testing.T) {
	fmt.Println("~~~~~~~~~~~~~~~Start greek factory test~~~~~~~~~~~~~~~")
	testList := []string{"north", "south", "others"}
	greekFactory := &GreekFactory{}
	for _, test := range testList {
		fmt.Println("======================")
		fmt.Printf("Creating %s pizza \n", test)
		pizza := greekFactory.GeneratePizza(test)
		if pizza != nil {
			pizza.Prepare()
			pizza.Bake()
			pizza.Cut()
			pizza.Box()
		} else {
			fmt.Println("Invalid pizza")
		}
	}

}
