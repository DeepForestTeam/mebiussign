package store

import (
	"github.com/asdine/storm"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"encoding/json"
	"sync"
)

type GlobalStore struct {
	mux       sync.Mutex
	db        *storm.DB
	TotalKeys int64
	storage   map[string]interface{}
}

var storage_instance GlobalStore

func init() {
	log.Info("* Init store")
	storage_instance.TotalKeys = 0
	storage_instance.storage = make(map[string]interface{})
}

func ConnectDB() (err error) {
	db_name, err := config.GetString("BOLT_DB")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Debug("Open BoltDB Storage")
	db, err := storm.Open(db_name, storm.AutoIncrement())
	if err != nil {
		log.Fatal(err)
		return
	}
	storage_instance.db = db
	return
}

func SaveObject(model_name, key string, object interface{}) (err error) {
	data, err := json.Marshal(object)
	storage_instance.mux.Lock()
	defer storage_instance.mux.Unlock()
	if err != nil {
		log.Error("Can not json encode object:", err)
		return
	}
	err = storage_instance.db.Set(model_name, key, data)
	if err != nil {
		log.Error("Can not save object:", err)
		return
	}
	log.Debug("Save:", key, "->", string(data))
	return
}
func GetObject(model_name, key string, object interface{}) (err error) {
	storage_instance.mux.Lock()
	defer storage_instance.mux.Unlock()
	var data []byte
	err = storage_instance.db.Get(model_name, key, &data)
	if err != nil {
		log.Error("Con not get object:", err)
		return
	}
	log.Debug("Load:", key, "->", string(data))
	err = json.Unmarshal(data, object)
	if err != nil {
		log.Error("Con not unjson object:", err)
	}
	return
}
func GetStat(model_name string) (err error) {

	//	GlobalStoreBarrel.db.Range()
	return
}
func CountObjects(model_name string) (count int, err error) {

	return
}