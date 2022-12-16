package facade

import "fmt"

type VoltageSensor struct {
}

func (v *VoltageSensor) GetMeasurement() error {
	fmt.Println("Voltage is ok")
	return nil
}
