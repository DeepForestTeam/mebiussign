package memstore

import "log"

type GlobalMemStore struct {
	TotalKeys int64
	storage   map[string]interface{}
}

var GlobalMemStoreInstance GlobalMemStore

func init() {
	log.Println("* Init store")
	GlobalMemStoreInstance.TotalKeys = 0
	GlobalMemStoreInstance.storage = make(map[string]interface{})
}

func (this *GlobalMemStore)GetByKey(key string) (value interface{}, ok bool) {
	return
}

func (this *GlobalMemStore)isKeyExist(key string) (ok bool) {
	return
}