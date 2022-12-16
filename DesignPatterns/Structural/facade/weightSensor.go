package facade

import "fmt"

type WeightSensor struct {
}

func (w *WeightSensor) GetMeasurement() error {
	fmt.Println("Weight is ok")
	return nil
}
