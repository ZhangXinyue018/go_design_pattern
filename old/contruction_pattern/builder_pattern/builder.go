package builder_pattern

import "fmt"

type LapTopBuilder interface {
	BuildCPU(str string) CPU

	BuildMonitor(str string) Monitor

	BuildDrive(str string) Drive
}

type FancyLapTopBuilder struct {
}

func (builder *FancyLapTopBuilder) BuildCPU(str string) CPU {
	fmt.Printf("Init CPU with name [%s] \n", str)
	return CPU{str}
}

func (builder *FancyLapTopBuilder) BuildMonitor(str string) Monitor {
	fmt.Printf("Init monitor with name [%s] \n", str)
	return Monitor{str}
}

func (builder *FancyLapTopBuilder) BuildDrive(str string) Drive {
	fmt.Printf("Init drive with name [%s] \n", str)
	return Drive{str}
}

type LapTopDirector interface {
	DirectLapTop() LapTop
}

type FancyLapTopDirector struct {
	builder FancyLapTopBuilder
}

func (director *FancyLapTopDirector) DirectLapTop() LapTop {
	CPU := director.builder.BuildCPU("Fancy CPU")
	Drive := director.builder.BuildDrive("Fancy Drive")
	Monitor := director.builder.BuildMonitor("Fancy Monitor")
	return LapTop{
		CPU:     CPU,
		Drive:   Drive,
		Monitor: Monitor,
	}
}
