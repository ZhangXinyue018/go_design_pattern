package abstract

type LaptopFactory interface {
	CreateCPU() CPU

	CreateDrive() Drive
}

type ExpensiveFactory struct {
}

func (factory *ExpensiveFactory) CreateCPU() CPU {
	return &ExpensiveCPU{}
}

func (factory *ExpensiveFactory) CreateDrive() Drive {
	return &ExpensiveDrive{}
}

type CheapFactory struct {
}

func (factory *CheapFactory) CreateCPU() CPU {
	return &CheapCPU{}
}

func (factory *CheapFactory) CreateDrive() Drive {
	return &CheapDrive{}
}
