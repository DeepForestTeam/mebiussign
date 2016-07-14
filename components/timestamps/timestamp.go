package timestamps

import (
	"time"
	"fmt"
	"crypto/sha512"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/store"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"errors"
	"encoding/hex"
	"encoding/binary"
	"bytes"
)

type TimeStampSignature struct {
	TimeZone      string    `json:"time_zone"`
	TimeStamp     string    `json:"time"`
	UnixTimeStamp int64     `json:"unix_time"`
	SaltHash      string    `json:"salt_hash"`
	PepperHash    string    `json:"pepper_hash"`
	MobiusTime    string    `json:"mobius_time"`
}

var lastState = TimeStampSignature{}

const (
	TimeStampStorage = "time_stamp"
)

func init() {
	log.Info("* Init timestamps")
}

var (
	ErrorInitError = errors.New("time init error")
	ErrorNotFound = errors.New("time stamp not found")
)

func (this *TimeStampSignature)GetCurrent() (err error) {
	time_now := time.Now().UTC()
	timestamp := time_now.Unix()
	if timestamp == lastState.UnixTimeStamp {
		copy(&lastState, this)
		return
	}
	this.UnixTimeStamp = timestamp
	this.TimeStamp = fmt.Sprintf("%d", timestamp)
	time_format, err := config.GetString("BASE_TIME_FORMAT")
	if err != nil {
		log.Critical("Can not get time format:", err)
		return
	}
	time_zone, err := config.GetString("BASE_TIME_ZONE")
	if err != nil {
		log.Critical("Can not get time zone:", err)
		return
	}

	this.TimeStamp = time_now.Format(time_format)
	this.UnixTimeStamp = timestamp
	this.TimeZone = time_zone
	err = this.sign()
	if err != nil {
		log.Critical("Can not sign time stamp:", err)
	}
	//Save last state
	copy(this, &lastState)
	err = store.Set(TimeStampStorage, this.MobiusTime, this)
	if err != nil {
		log.Critical("Can not store time stamp:", err)
	}
	return
}

func (this *TimeStampSignature)GetBySign(hash string) (err error) {
	err = store.Get(TimeStampStorage, hash, this)
	if err != nil {
		log.Error("Can not get timestamp from storage:", err)
		return ErrorNotFound
	}
	return nil
}
func (this *TimeStampSignature)sign() (err error) {
	base_salt_string, err := config.GetString("BASE_TIME_SALT")
	if err != nil {
		log.Error("Can not get base time salt:", err)
		return
	}
	first := false
	var bytes_block []byte
	last_block := TimeStampSignature{}
	_, err = store.Last(TimeStampStorage, &last_block)
	if err != nil && err != store.ErrKeyNotFound && err != store.ErrSectionNotFound {
		log.Error("Can not get last block:", err)
		return
	}
	//empty MobiusTime tape. Init
	if err == store.ErrSectionNotFound || err == store.ErrKeyNotFound {
		log.Warning("First Time signature!")
		this.SaltHash = base_salt_string
		err = nil
		first = true
	} else {
		this.SaltHash = last_block.MobiusTime
	}
	base_salt, _ := hex.DecodeString(base_salt_string)
	time_stamp := intToBytes(this.UnixTimeStamp)
	salt_hash, _ := hex.DecodeString(this.SaltHash)
	pepper := []byte("Pepper")

	bytes_block = append(bytes_block, salt_hash...)
	bytes_block = append(bytes_block, pepper...)
	bytes_block = append(bytes_block, base_salt...)
	bytes_block = append(bytes_block, time_stamp...)

	log.Debug("DataBlock len:", len(bytes_block))
	hasher := sha512.New512_256()
	hasher.Write(bytes_block)
	signature := hasher.Sum(nil)
	this.PepperHash = fmt.Sprintf("%X", pepper)
	this.MobiusTime = fmt.Sprintf("%X", signature)
	return
}
func copy(from *TimeStampSignature, to *TimeStampSignature) {
	to.TimeZone = from.TimeZone
	to.TimeStamp = from.TimeStamp
	to.UnixTimeStamp = from.UnixTimeStamp
	to.PepperHash = from.PepperHash
	to.SaltHash = from.SaltHash
	to.MobiusTime = from.MobiusTime
}

func intToBytes(num int64) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}