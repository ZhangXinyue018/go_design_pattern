package example

import "fmt"

type IController interface {
	DoControl()
}

type IAdapter interface {
	CheckController(controller IController) bool

	DoAdapt(controller IController)
}

// Simple
type SimpleController struct {
}

func (c *SimpleController) DoControl() {
	fmt.Println("Simple controller doing control")
}

type SimpleAdapter struct {
}

func (a *SimpleAdapter) CheckController(controller IController) bool {
	_, ok := controller.(*SimpleController)
	return ok
}

func (a *SimpleAdapter) DoAdapt(controller IController) {
	if !a.CheckController(controller) {
		fmt.Println("controller type invalid")
	} else {
		controller.DoControl()
		fmt.Println("adapt simple controller")
	}
}

// Complex
type ComplexController struct {
}

func (c *ComplexController) DoControl() {
	fmt.Println("Complex controller doing control")
}

type ComplexAdapter struct {
}

func (a *ComplexAdapter) CheckController(controller IController) bool {
	_, ok := controller.(*ComplexController)
	return ok
}

func (a *ComplexAdapter) DoAdapt(controller IController) {
	if !a.CheckController(controller) {
		fmt.Println("controller type invalid")
	} else {
		controller.DoControl()
		fmt.Println("adapt complex controller")
	}
}

type Holder struct {
	Records []IAdapter
}

func (holder *Holder) AddAdapter(adapter IAdapter) {
	holder.Records = append(holder.Records, adapter)
}

func (holder *Holder) DoHold(controller IController) {
	for _, adapter := range holder.Records {
		if adapter.CheckController(controller) {
			adapter.DoAdapt(controller)
			return
		}
	}
}
