package example

import "testing"

func TestAdapter(t *testing.T) {
	holder := Holder{}
	holder.AddAdapter(&SimpleAdapter{})
	holder.AddAdapter(&ComplexAdapter{})
	//holder.DoHold(&SimpleController{})
	holder.DoHold(&ComplexController{})
}
