package abstract

import "fmt"

type CPU interface {
	Calculate() ()
}

type ExpensiveCPU struct {
}

func (cpu *ExpensiveCPU) Calculate() () {
	fmt.Println("Calculate fast!!!")
}

type CheapCPU struct {
}

func (cpu *CheapCPU) Calculate() () {
	fmt.Println("Calculate slow!!!")
}

type Drive interface {
	Store() ()
}

type ExpensiveDrive struct {
}

func (drive *ExpensiveDrive) Store() () {
	fmt.Println("Store more!!!")
}

type CheapDrive struct {
}

func (drive *CheapDrive) Store() () {
	fmt.Println("Store less!!!")
}
