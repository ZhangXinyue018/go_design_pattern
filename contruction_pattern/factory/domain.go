package factory

import "fmt"

type Animal interface {
	Run()
}

type Dog struct {
	Name string
	Animal
}

func (dog *Dog) Run() {
	fmt.Printf("dog %s is running!!! \n", dog.Name)
}

type Cat struct {
	Name string
	Animal
}

func (cat *Cat) Run() {
	fmt.Printf("cat %s is running!!!\n", cat.Name)
}
