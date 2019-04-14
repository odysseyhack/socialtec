package store

import "log"

// Store interface for all storage
type Store interface {
	Get(string, interface{}) error
	Set(string, interface{}) error
}

var Default Store

func NewStore(filename string) Store {
	db, err := newDB(filename)
	if err != nil {
		log.Fatalf("Failed initializing storage. Filename:%s Err:%+v", filename, err)
	}
	return db
}
