package facade

import "fmt"

type WaterSensor struct {
}

func (w *WaterSensor) GetMeasurement() error {
	fmt.Println("Water is ok")
	return nil
}
