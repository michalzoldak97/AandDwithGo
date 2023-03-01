package factory

type Truck3 struct {
	Car
}

func NewTruck3() Vehicle {
	return &Truck3{
		Car: Car{
			model:  "Truck3",
			gears:  12,
			hp:     400,
			wheels: 6,
		},
	}
}
