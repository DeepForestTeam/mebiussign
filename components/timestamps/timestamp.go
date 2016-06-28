package timestamps

type TimeStampSignature struct {
	UnixTimeStamp int64   `json:"time_stamp"`
	TimeHashSign  string  `json:"time_hash"`
}

type TimeLine struct {
	TimeStampIndex   map[string]int64
	CurrenttimeStamp TimeStampSignature
	TimeLine        []TimeStampSignature
}

var MasterTimeSalt string

func init(){
	MasterTimeSalt="READ_FROM_CONFIG"
}

func (this *TimeStampSignature)GetCurrent(){

}