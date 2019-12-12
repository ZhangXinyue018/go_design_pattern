package builder

import "github.com/go_design_pattern/creational/builder/domain"

type Director struct {
	Builder domain.HouseBuilder
}

func (director *Director) Build() *domain.House {
	director.Builder.NewHouse()
	director.Builder.BuildBasic()
	director.Builder.BuildWall()
	director.Builder.BuildTop()
	return director.Builder.GetResult().(*domain.House)
}
