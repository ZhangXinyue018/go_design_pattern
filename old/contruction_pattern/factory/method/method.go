package method

import "github.com/go_design_pattern/old/contruction_pattern/factory"

type AnimalFactory interface {
	CreateAnimal() factory.Animal
}

type CatFactory struct {
}

func (catFactory *CatFactory) CreateAnimal() factory.Animal {
	return &factory.Cat{Name: "catty"}
}

type DogFactory struct {
}

func (dogFactory *DogFactory) CreateAnimal() factory.Animal {
	return &factory.Dog{Name: "doggy"}
}
