package simple

import "go_design_pattern/contruction_pattern/factory"

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
