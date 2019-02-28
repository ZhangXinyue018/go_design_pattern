package builder_pattern

type LapTop struct {
	CPU     CPU
	Drive   Drive
	Monitor Monitor
}

type CPU struct {
	Name string
}

type Drive struct {
	Name string
}

type Monitor struct {
	Name string
}
