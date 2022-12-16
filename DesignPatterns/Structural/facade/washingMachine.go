package facade

type WashingMachine struct {
	waterPump Pump
	wmMotor   Motor
	sensors   []Sensor
}

func NewWashingMachine() (*WashingMachine, error) {
	washingMachine := &WashingMachine{
		waterPump: &WaterPump{},
		wmMotor:   &WMMotor{},
	}
	vSensor := &VoltageSensor{}
	wrSensor := &WaterSensor{}
	wtSensor := &WeightSensor{}
	washingMachine.sensors = []Sensor{vSensor, wrSensor, wtSensor}
	return washingMachine, nil
}

func (wm *WashingMachine) CheckSensors() error {
	for _, s := range wm.sensors {
		err := s.GetMeasurement()
		if err != nil {
			return err
		}
	}

	return nil
}

func (wm *WashingMachine) StartWashingMachine() error {
	err := wm.wmMotor.StartMotor()
	if err != nil {
		return err
	}

	err = wm.waterPump.StartPump()
	if err != nil {
		return err
	}

	err = wm.CheckSensors()
	if err != nil {
		return err
	}

	return nil
}

func (wm *WashingMachine) StopWashingMachine() error {
	err := wm.wmMotor.StopMotor()
	if err != nil {
		return err
	}

	err = wm.waterPump.StopPump()
	if err != nil {
		return err
	}

	err = wm.CheckSensors()
	if err != nil {
		return err
	}

	return nil
}
