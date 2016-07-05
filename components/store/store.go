package store

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GlobalMemStore struct {
	//db
	TotalKeys int64
	storage   map[string]interface{}
}

type TimeStamp struct {
	gorm.Model
	ID       int64  //Field name: id, type:bigint(20)
	UnixTime int64  //Field name: unix_time, type:bigint(20)
	TimeSign string //Field name: time_sign, type:varchar(128)
}

var GlobalMemStoreInstance GlobalMemStore

func init() {
	log.Println("* Init store")
	GlobalMemStoreInstance.TotalKeys = 0
	GlobalMemStoreInstance.storage = make(map[string]interface{})
}

func (this *GlobalMemStore)ConnectDB() {

}

func (this *GlobalMemStore)GetByKey(key string) (value interface{}, ok bool) {
	return
}

func (this *GlobalMemStore)isKeyExist(key string) (ok bool) {
	return
}