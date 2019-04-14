package store

import (
	"encoding/json"
	"errors"

	"github.com/tidwall/buntdb"
)

// DB is keyvalue store
type DB struct {
	*buntdb.DB
}

var ErrorNotFound = errors.New("Not Found")

func newDB(storeFile string) (*DB, error) {
	client, err := buntdb.Open(storeFile)
	if err != nil {
		return nil, err
	}
	return &DB{DB: client}, nil
}

// Get gets value given id
func (db *DB) Get(id string, val interface{}) error {
	var resultStr string
	var err error
	db.DB.View(func(tx *buntdb.Tx) error {
		resultStr, err = tx.Get(id, true)
		return nil
	})

	if err != nil {
		if err == buntdb.ErrNotFound {
			return ErrorNotFound
		}
		return err
	}
	if err := json.Unmarshal([]byte(resultStr), val); err != nil {
		return err
	}
	return nil
}

// Set sets a value in db
func (db *DB) Set(key string, val interface{}) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, string(data), nil)
		return err
	})
}
