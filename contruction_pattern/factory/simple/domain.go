package simple

import "fmt"

type Animal interface {
	Run()
}

type Dog struct {
	Name string
	Animal
}

func (dog *Dog) Run() {
	fmt.Println("dog is running!!!")
}

type Cat struct {
	Name string
	Animal
}

func (cat *Cat) Run() {
	fmt.Println("cat is running!!!")
}
