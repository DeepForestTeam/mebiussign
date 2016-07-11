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

type LastState struct {
	LastTimeStamp TimeStampSignature
	TotalRecords  int64
}

const (
	TimeStampStorage = "time_stamp"
	TimeStampStorageState = "time_stamp_state"
)

func init() {
	log.Info("* Init timestamps")
}

var ObjectModelName = "time_stamp"

var (
	ErrorNotFound = errors.New("Time stamp not found")
)

func (this *TimeStampSignature)GetCurrent() (err error) {
	hash := ""
	time_now := time.Now().UTC()
	timestamp := time_now.Unix()
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
	hash = fmt.Sprintf("%X", sha512.Sum512_256([]byte(signed_value)))
	this.TimeHashSign = hash
	this.TimeStamp = time_now.Format(time_format)
	this.UnixTimeStamp = timestamp
	this.TimeZone = "UTC"
	store.SaveObject(TimeStampStorage, hash, this)
	return
}

func (this *TimeStampSignature)GetBySign(hash string) (err error) {
	err = store.GetObject(TimeStampStorage, hash, this)
	return
}