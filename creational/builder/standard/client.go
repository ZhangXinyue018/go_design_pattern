package standard

type Director struct {
	Builder HouseBuilder
}

func (director *Director) Build() *House {
	director.Builder.NewHouse()
	director.Builder.BuildBasic()
	director.Builder.BuildWall()
	director.Builder.BuildTop()
	return director.Builder.GetResult().(*House)
}
