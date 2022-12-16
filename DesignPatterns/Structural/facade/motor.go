package facade

type Motor interface {
	StartMotor() error
	StopMotor() error
}
