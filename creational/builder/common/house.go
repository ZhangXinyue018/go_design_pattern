package common

type House struct {
	Basic string
	Wall  string
	Top   string
}

type HouseBuilder interface {
	BuildBasic() HouseBuilder
	BuildWall() HouseBuilder
	BuildTop() HouseBuilder
	GetResult() interface{}
}

type CommonHouseBuilder struct {
	House *House
}

func (builder *CommonHouseBuilder) BuildBasic() HouseBuilder {
	builder.House.Basic = "common house basic"
	return builder
}

func (builder *CommonHouseBuilder) BuildWall() HouseBuilder {
	builder.House.Wall = "common house wall"
	return builder
}

func (builder *CommonHouseBuilder) BuildTop() HouseBuilder {
	builder.House.Top = "common house top"
	return builder
}

func (builder *CommonHouseBuilder) GetResult() interface{} {
	return builder.House
}

func NewCommonHouseBuilder() *CommonHouseBuilder {
	return &CommonHouseBuilder{House: &House{}}
}
