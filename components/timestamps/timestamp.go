package timestamps

import (
	"time"
	"fmt"
	"crypto/sha512"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/store"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

type TimeStampSignature struct {
	store.StoreObject
	UnixTimeStamp int64   `json:"time_stamp"`
	TimeHashSign  string  `json:"time_hash"`
}

type TimeLine struct {
	TimeStampIndex map[string]int64
	LastTimeStamp  TimeStampSignature
	TimeLine       []TimeStampSignature
}

var MasterTimeLine TimeLine

func init() {
	log.Info("* Init timestamps")
	MasterTimeLine.TimeStampIndex = make(map[string]int64)
}

func (this *TimeStampSignature)GetCurrent() (hash string, err error) {
	this.ObjectModelName = "time_stamp"
	timestamp := time.Now().Unix()
	time_stamp_string := fmt.Sprintf("%d", timestamp)
	time_base_salt, err := config.GlobalConfig.GetString("BASE_TIME_SALT")
	if err != nil {
		return
	}
	signed_value := time_base_salt + time_stamp_string
	hash = fmt.Sprintf("%X", sha512.Sum512_256([]byte(signed_value)))
	this.TimeHashSign = hash
	this.UnixTimeStamp = timestamp
	this.ObjectUID = hash
	this.Save()
	return
}
