package memstore

var GlobalMemStore struct {
	TotalKeys int64
	storage   map[string]interface{}
}

func init() {
	GlobalMemStore.storage = make(map[string]interface{})
}

