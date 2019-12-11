package prototype

type CopyObject interface {
	Clone() CopyObject
}

type Sheep struct {
	Name string
	Age  int
}

// simple clone
func (sheep *Sheep) Clone() CopyObject {
	obj := *sheep
	return &obj
}
