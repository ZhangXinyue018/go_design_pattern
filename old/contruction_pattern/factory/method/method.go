package method

import "go_design_pattern/contruction_pattern/factory"

type AnimalFactory interface {
	CreateAnimal() (factory.Animal)
}

type CatFactory struct {
}

func (catFactory *CatFactory) CreateAnimal() (factory.Animal) {
	return &factory.Cat{Name: "catty"}
}

type DogFactory struct {
}

func (dogFactory *DogFactory) CreateAnimal() (factory.Animal) {
	return &factory.Dog{Name: "doggy"}
}
