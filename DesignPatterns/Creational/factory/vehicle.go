package factory

type Vehicle interface {
	Model() string
	NumGears() int
	NumHP() int
	NumWheels() int
}
