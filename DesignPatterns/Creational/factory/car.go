package factory

type Car struct {
	model  string
	gears  int
	hp     int
	wheels int
}

func (c *Car) Model() string {
	return c.model
}

func (c *Car) NumGears() int {
	return c.gears
}

func (c *Car) NumHP() int {
	return c.hp
}

func (c *Car) NumWheels() int {
	return c.wheels
}
