package controls


import (
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/timestamps"
)

type TimeApiController struct {
	forest.Control
}

func (this *TimeApiController)Get() {
	defer this.ServeJSON()
	ts := timestamps.TimeStampSignature{}
	time_hash := this.Context.UrlVars["time_hash"]
	if time_hash == "" {
		err := ts.GetCurrent()
		if err != nil {
			log.Error("Can not create new time stamp!")
			this.Data=ErrorMessage{Result:"Server error", ResultCode:500}
			return
		}
	} else {
		//@todo Check hash
		err := ts.GetBySign(time_hash)
		if err != nil {
			this.Data=ErrorMessage{Result:"Hash not found", ResultCode:404}
			return
		}
	}
	this.Data = ts
}
