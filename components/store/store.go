package store

import (
	"github.com/DeepForestTeam/mobiussign/components/log"
)

type GlobalStore struct {
	driver StorageDriver
}

type StorageDriver interface {
	Connect() error
	Close()
	Set(section, key string, data interface{}) error
	Get(section, key string, data interface{}) error
	Count(section string) (int, error)
	Last(section string, data interface{}) (string, error)
}

var storage GlobalStore

func init() {
	log.Info("* Init store")
	storage.driver = &BoltDriver{}
}

func ConnectDB() (err error) {
	err = storage.driver.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Debug("Connect Driver Storage")
	return
}

func Set(model_name, key string, object interface{}) (err error) {
	return storage.driver.Set(model_name, key, object)
}
func Get(model_name, key string, object interface{}) (err error) {
	return storage.driver.Get(model_name, key, object)
}
func Count(model_name string) (count int, err error) {
	return storage.driver.Count(model_name)
}