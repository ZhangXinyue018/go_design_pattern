package class

func ChargePhone() {
	phone := Phone{Voltage5V: &VoltageAdapter{}}
	phone.Charge()
}
