package domain

type House struct {
	Basic string
	Wall  string
	Top   string
}

type HouseBuilder interface {
	NewHouse()
	BuildBasic()
	BuildWall()
	BuildTop()
	GetResult() interface{}
}

type CommonHouseBuilder struct {
	House *House
}

func (builder *CommonHouseBuilder) NewHouse() {
	builder.House = &House{}
}

func (builder *CommonHouseBuilder) BuildBasic() {
	builder.House.Basic = "common house basic"
}

func (builder *CommonHouseBuilder) BuildWall() {
	builder.House.Wall = "common house wall"
}

func (builder *CommonHouseBuilder) BuildTop() {
	builder.House.Top = "common house top"
}

func (builder *CommonHouseBuilder) GetResult() interface{} {
	return builder.House
}
