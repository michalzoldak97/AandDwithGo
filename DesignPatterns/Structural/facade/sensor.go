package facade

type Sensor interface {
	GetMeasurement() error
}
