package facade

import "fmt"

type WaterPump struct {
}

func (wp *WaterPump) StartPump() error {
	fmt.Println("Water pump launched and OK")
	return nil
}

func (wp *WaterPump) StopPump() error {
	fmt.Println("Water pump stopped and OK")
	return nil
}
