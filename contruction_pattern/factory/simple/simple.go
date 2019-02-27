package simple

func AnimalFactory(animalType string) Animal {
	switch animalType {
	case "dog":
		return &Dog{Name: "doggy"}
	case "cat":
		return &Cat{Name: "catty"}
	default:
		return nil
	}

}
