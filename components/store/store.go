package store

import (

	"github.com/asdine/storm"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/config"
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

type StoreObject struct {
	ID              int64
	ObjectModelName string
	ObjectUID       string
	ObjectID        int64
	ObjectShardID   int64
	ObjectOwnerUID  string
}

func (this *StoreObject)Init() {

}
func (this *StoreObject)Save() (err error) {
	err = nil
	err = GlobalStoreBarrel.db.Set(this.ObjectModelName, this.ObjectUID, this)
	//log.Fatalln(err)
	return
}
func (this *StoreObject)Get() (err error) {
	err = nil
	err = GlobalStoreBarrel.db.Get(this.ObjectModelName, this.ObjectUID, this)
	return
}