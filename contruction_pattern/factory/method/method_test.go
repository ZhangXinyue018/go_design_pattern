package method

import (
	"testing"
)

func TestAnimalFactory(t *testing.T) {
	t.Run("general test", func(t *testing.T) {
		got1 := (&CatFactory{}).CreateAnimal()
		got1.Run()
		got2 := (&DogFactory{}).CreateAnimal()
		got2.Run()
	})
}
