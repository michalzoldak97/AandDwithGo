package decorator

type Car struct {
	price int
}

func (c *Car) SetPrice(price int) {
	c.price = price
}

func (c *Car) GetOffer() int {
	return c.price
}
