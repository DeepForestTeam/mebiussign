package store

import "github.com/boltdb/bolt"

type BoltWarper struct {
	db        *bolt.DB
	connected bool
}

type StorageObject struct {

}

type StorageDriver interface {
	Connect() error
	Set(section, key string, data *interface{}) error
	Get(section, key string, data *interface{}) error
	Count(section string) (int, error)
	Last(section string, data *interface{}) error
}

func init() {

}
