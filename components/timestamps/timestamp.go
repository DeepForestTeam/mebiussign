package timestamps

import (
	"log"
	"time"
	"fmt"
	"crypto/sha512"
	"github.com/DeepForestTeam/mobiussign/components/config"
)

type TimeStampSignature struct {
	UnixTimeStamp int64   `json:"time_stamp"`
	TimeHashSign  string  `json:"time_hash"`
}

type TimeLine struct {
	TimeStampIndex   map[string]int64
	CurrenttimeStamp TimeStampSignature
	TimeLine         []TimeStampSignature
}

var MasterTimeSalt string

func init() {
	log.Println("* Init timestamps")
}

func (this *TimeStampSignature)GetCurrent() (hash string, err error) {
	timestamp := time.Now().UnixNano()
	log.Println("Tilestmp:", timestamp)
	time_stamp_string := fmt.Sprintf("%d", timestamp)
	time_base_hash, err := config.GlobalConfig.GetString("BASE_TIME_HASH")
	if err != nil {
		return
	}
	signed_value := time_base_hash + time_stamp_string
	hash = string(sha512.Sum512_256([]byte(signed_value)))
	return
}

func TimeTick() {

}