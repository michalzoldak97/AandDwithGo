package facade

import "fmt"

type WMMotor struct {
}

func (wm *WMMotor) StartMotor() error {
	fmt.Println("Motor launched and OK")
	return nil
}

func (wm *WMMotor) StopMotor() error {
	fmt.Println("Motor stopped and OK")
	return nil
}
