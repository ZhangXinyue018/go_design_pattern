package common

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	house1 := NewCommonHouseBuilder().BuildBasic().BuildWall().BuildTop().GetResult().(*House)
	house2 := NewCommonHouseBuilder().BuildBasic().BuildWall().BuildTop().GetResult().(*House)
	fmt.Printf("house1 of %v and addr %p\n", *house1, house1)
	fmt.Printf("house2 of %v and addr %p\n", *house2, house2)
}
