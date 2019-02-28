package simple

import (
	"fmt"
	"go_design_pattern/contruction_pattern/factory"
	"testing"
)

func TestGenAnimal(t *testing.T) {
	type args struct {
		animalType string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"test1", args{"dog"}, &factory.Dog{}},
		{"test2", args{"cat"}, &factory.Cat{}},
		{"test3", args{"not exist"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnimalFactory(tt.args.animalType)
			fmt.Printf("[%s]: got value %v \n", tt.name, got)
			if got != nil {
				got.Run()
			}
		})
	}
}
