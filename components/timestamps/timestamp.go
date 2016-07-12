package timestamps

import (
	"time"
	"fmt"
	"crypto/sha512"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/store"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"errors"
)

type TimeStampSignature struct {
	TimeZone      string    `json:"time_zone"`
	TimeStamp     string    `json:"time"`
	UnixTimeStamp int64     `json:"unix_time"`
	TimeHashSign  string    `json:"time_hash"`
}

var lastState = TimeStampSignature{}

const (
	TimeStampStorage = "time_stamp"
)

func init() {
	log.Info("* Init timestamps")
}

var (
	ErrorNotFound = errors.New("Time stamp not found")
)

func (this *TimeStampSignature)GetCurrent() (err error) {

	time_now := time.Now().UTC()
	timestamp := time_now.Unix()
	if timestamp == lastState.UnixTimeStamp {
		this.TimeZone = "UTC"
		this.UnixTimeStamp = lastState.UnixTimeStamp
		this.TimeHashSign = lastState.TimeHashSign
		this.TimeStamp = lastState.TimeStamp
		return
	}
	time_stamp_string := fmt.Sprintf("%d", timestamp)
	time_base_salt, err := config.GetString("BASE_TIME_SALT")
	if err != nil {
		return
	}
	time_format, err := config.GetString("BASE_TIME_FORMAT")
	if err != nil {
		return
	}
	signed_value := time_base_salt + time_stamp_string
	hash := ""
	hash = fmt.Sprintf("%X", sha512.Sum512_256([]byte(signed_value)))
	this.TimeHashSign = hash
	this.TimeStamp = time_now.Format(time_format)
	this.UnixTimeStamp = timestamp
	this.TimeZone = "UTC"
	//Save last state
	lastState.TimeZone = "UTC"
	lastState.UnixTimeStamp = this.UnixTimeStamp
	lastState.TimeHashSign = this.TimeHashSign
	lastState.TimeStamp = this.TimeStamp
	store.Set(TimeStampStorage, hash, this)
	return
}

func (this *TimeStampSignature)GetBySign(hash string) (err error) {
	err = store.Get(TimeStampStorage, hash, this)
	if err == store.ErrKeyNotFound {
		return ErrorNotFound
	}
	return nil
}