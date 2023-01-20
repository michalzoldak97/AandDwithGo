package proxy

type DB interface {
	RunQuery(string) (int, error)
}
