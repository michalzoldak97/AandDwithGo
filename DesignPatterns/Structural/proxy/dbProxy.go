package proxy

import "errors"

type DBProxy struct {
	db *RealDB
}

func NewDBServer(limit int) *DBProxy {
	p := &DBProxy{}
	p.db = NewDBConn(limit)
	return p
}

func (p *DBProxy) RunQuery(q string) (int, error) {
	if !p.db.IsAvailable() {
		return -1, errors.New("no avaliable DB connections")
	}

	res, err := p.db.RunQuery(q)
	if err != nil {
		return -1, err
	}

	return res, nil
}
