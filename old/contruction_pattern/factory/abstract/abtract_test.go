package abstract

import (
	"testing"
)

func TestAbstrctFactory(t *testing.T) {
	t.Run("general test", func(t *testing.T) {
		laptopFactory1 := &ExpensiveFactory{}
		cpu1 := laptopFactory1.CreateCPU()
		cpu1.Calculate()
		drive1 := laptopFactory1.CreateDrive()
		drive1.Store()

		laptopFactory2 := &CheapFactory{}
		cpu2 := laptopFactory2.CreateCPU()
		cpu2.Calculate()
		drive2 := laptopFactory2.CreateDrive()
		drive2.Store()
	})

}
