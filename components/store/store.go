package store

import (
	"github.com/asdine/storm"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"encoding/json"
)

type GlobalStore struct {
	db        *storm.DB
	TotalKeys int64
	storage   map[string]interface{}
}

var GlobalStoreBarrel GlobalStore

func init() {
	log.Info("* Init store")
	GlobalStoreBarrel.TotalKeys = 0
	GlobalStoreBarrel.storage = make(map[string]interface{})
}

func (this *GlobalStore)ConnectDB() (err error) {
	db_name, err := config.GlobalConfig.GetString("BOLT_DB")
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
	this.db = db
	return
}

func SaveObject(model_name, key string, object interface{}) (err error) {
	data, err := json.Marshal(object)
	if err != nil {
		log.Error("Can not json encode object:", err)
		return
	}
	err = GlobalStoreBarrel.db.Set(model_name, key, data)
	if err != nil {
		log.Error("Can not save object:", err)
		return
	}
	log.Debug("Save:", key, "->", string(data))
	return
}
func GetObject(model_name, key string, object interface{}) (err error) {
	var data []byte
	err = GlobalStoreBarrel.db.Get(model_name, key, &data)
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