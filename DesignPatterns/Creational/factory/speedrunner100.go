package factory

type Speedrunner100 struct {
	Car
}

func NewSpeedrunner100() Vehicle {
	return &Speedrunner100{
		Car: Car{
			model:  "Speedrunner100",
			gears:  5,
			hp:     320,
			wheels: 4,
		},
	}
}
