package facade

type Pump interface {
	StartPump() error
	StopPump() error
}
