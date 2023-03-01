package factory

import "errors"

func GetCar(model string) (Vehicle, error) {
	if model == "Speedrunner100" {
		return NewSpeedrunner100(), nil
	} else if model == "Truck3" {
		return NewTruck3(), nil
	} else {
		return &Car{}, errors.New("Model not found")
	}
}
