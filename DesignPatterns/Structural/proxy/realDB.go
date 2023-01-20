package proxy

import (
	"errors"
	"math/rand"
	"time"
)

type RealDB struct {
	maxDBLoad int
	dbLoad    int
}

func NewDBConn(limit int) *RealDB {
	return &RealDB{
		maxDBLoad: limit,
		dbLoad:    0,
	}
}

func (db *RealDB) IsAvailable() bool {
	return db.dbLoad < db.maxDBLoad
}

func (db *RealDB) addLoad(change int) error {
	load := db.dbLoad + change

	if load > db.maxDBLoad {
		return errors.New("DB load exceeded, refusing connection")
	}

	db.dbLoad = load
	return nil
}

func (db *RealDB) RunQuery(q string) (int, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rows := r1.Intn(100)

	err := db.addLoad(rows)
	if err != nil {
		return -1, err
	}

	return rows, nil
}
