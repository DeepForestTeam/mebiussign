package store

import (
	"reflect"
	"github.com/boltdb/bolt"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/store/codecs/gob"
	"github.com/DeepForestTeam/mobiussign/components/store/codecs/json"
	"github.com/DeepForestTeam/mobiussign/components/store/codecs/bson"
	"sync"
	"bytes"
	"encoding/binary"
)

type EncodeDecoder interface {
	Encode(v interface{}) ([]byte, error)
	Decode(b []byte, v interface{}) error
}

type BoltDriver struct {
	db        *bolt.DB
	connected bool
	Codec     EncodeDecoder
	mux       map[string]*sync.Mutex
}

type BoltIndex struct {
	ID  uint64
	Key string
}

const (
	indexPostfix = "_index"
)

func (this *BoltDriver)Connect() (err error) {
	db_name, err := config.GetString("BOLT_DB")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Debug("BoltDriver/Warper: Open BoltDB Storage, db file path:", db_name)
	this.db, err = bolt.Open(db_name, 0600, nil)
	if err != nil {
		log.Error("Can notopen BoltDB file:", err)
	}
	this.mux = make(map[string]*sync.Mutex)
	this.setDefaultCodec()
	return
}
func (this *BoltDriver)Close() {
	this.lockAll()
	defer this.unlockAll()
	this.db.Close()
}

func (this *BoltDriver)Set(bucket_name, key string, data interface{}) (err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	ref := reflect.ValueOf(data)
	if !ref.IsValid() || ref.Kind() != reflect.Ptr {
		return ErrPtrNeeded
	}
	err = this.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
		if err != nil {
			return err
		}
		value, err := this.Codec.Encode(data)
		if err != nil {
			return err
		}
		num, err := bucket.NextSequence()
		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}
		//Crete Index
		index_bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name + indexPostfix))
		if err != nil {
			return err
		}
		id := uintToBytes(num)
		err = index_bucket.Put(id, []byte(key))
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (this *BoltDriver)Get(bucket_name, key string, data interface{}) (err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	ref := reflect.ValueOf(data)
	if !ref.IsValid() || ref.Kind() != reflect.Ptr {
		return ErrPtrNeeded
	}
	err = this.db.View(func(tx *bolt.Tx) error {
		bucket, err := this.getBucket(tx, bucket_name)
		if err != nil {
			return err
		}
		val := bucket.Get([]byte(key))
		if len(val) == 0 {
			return ErrKeyNotFound
		}
		err = this.Codec.Decode(val, data)
		return err
	})
	return
}

func (this *BoltDriver)Count(bucket_name string) (count int, err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	err = this.db.View(func(tx *bolt.Tx) error {
		bucket, err := this.getBucket(tx, bucket_name)
		if err != nil {
			return err
		}
		stats := bucket.Stats()
		count = stats.KeyN
		return err
	})
	return
}

func (this *BoltDriver)Last(bucket_name string, data interface{}) (key string, err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	ref := reflect.ValueOf(data)
	if !ref.IsValid() || ref.Kind() != reflect.Ptr {
		return key, ErrPtrNeeded
	}
	err = this.db.View(func(tx *bolt.Tx) error {
		index_bucket, err := this.getBucket(tx, bucket_name + indexPostfix)
		if err != nil {
			return err
		}
		_, index_val := index_bucket.Cursor().Last()
		if index_val == nil {
			return ErrNotIndexed
		}
		key = string(index_val)
		if err != nil {
			return err
		}
		bucket, err := this.getBucket(tx, bucket_name)
		val := bucket.Get([]byte(key))
		if len(val) == 0 {
			return ErrKeyNotFound
		}
		err = this.Codec.Decode(val, data)
		return err
	})

	return
}
func (this *BoltDriver)getBucket(tx *bolt.Tx, bucket_name string) (bucket *bolt.Bucket, err error) {
	bucket = tx.Bucket([]byte(bucket_name))
	if bucket == nil {
		log.Error("BoltDriver/Warper: Bucket not found:", bucket_name)
		return bucket, ErrSectionNotFound
	}
	return
}

func (this *BoltDriver)setCodec(name string) {
	switch name {
	case "json":
		this.Codec = json.Codec
	case "gob":
		this.Codec = gob.Codec
	case "bson":
		this.Codec = bson.Codec
	}
}
func (this *BoltDriver)setDefaultCodec() {
	this.setCodec("json")
}
func (this *BoltDriver)lockBucket(bucket_name string) {
	if mux, ok := this.mux[bucket_name]; ok {
		mux.Lock()
	} else {
		mux = new(sync.Mutex)
		mux.Lock()
		this.mux[bucket_name] = mux
	}
}
func (this *BoltDriver)unlockBucket(bucket_name string) {
	if mux, ok := this.mux[bucket_name]; ok {
		mux.Unlock()
	}
}
func (this *BoltDriver)lockAll() {
	if len(this.mux) != 0 {
		for index, _ := range this.mux {
			this.mux[index].Lock()
		}
	}
}
func (this *BoltDriver)unlockAll() {
	if len(this.mux) != 0 {
		for index, _ := range this.mux {
			this.mux[index].Unlock()
		}
	}
}

func init() {

}

func uintToBytes(num uint64) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}


