package builder

import (
	"fmt"
	"github.com/go_design_pattern/creational/builder/domain"
	"testing"
)

func TestClient(t *testing.T) {
	director := Director{Builder: &domain.CommonHouseBuilder{}}
	house1 := director.Build()
	house2 := director.Build()
	fmt.Printf("house1 of %v and addr %p\n", *house1, house1)
	fmt.Printf("house2 of %v and addr %p\n", *house2, house2)
}

func TestTemp(t *testing.T){}