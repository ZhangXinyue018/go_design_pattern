package class

import "fmt"

type Voltage5V interface {
	GetOutput()
}

type Voltage220V struct {
}

func (v *Voltage220V) Output220() {
	fmt.Println("output 220")
}

type VoltageAdapter struct {
	Voltage220V
}

func (adapter *VoltageAdapter) GetOutput() {
	fmt.Println("get 220 result")
	adapter.Voltage220V.Output220()
	fmt.Println("transfer to 5v")
}

type Phone struct {
	Voltage5V Voltage5V
}

func (phone *Phone) Charge() {
	phone.Voltage5V.GetOutput()
}
