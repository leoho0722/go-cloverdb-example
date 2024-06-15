package database

import (
	"sync"

	cloverv2 "github.com/ostafen/clover/v2"
)

type Database struct {
	Context *cloverv2.DB
	Mutex   sync.Mutex
}

var DB *Database

// ConnectDB 用來連結 CloverDB
func ConnectDB() error {
	const (
		dbDir = "./"
	)
	db, err := cloverv2.Open(dbDir)
	if err != nil {
		return err
	}
	DB = &Database{Context: db}
	return nil
}
