package simple

import "github.com/go_design_pattern/old/contruction_pattern/factory"

func AnimalFactory(animalType string) factory.Animal {
	switch animalType {
	case "dog":
		return &factory.Dog{Name: "doggy"}
	case "cat":
		return &factory.Cat{Name: "catty"}
	default:
		return nil
	}

}
