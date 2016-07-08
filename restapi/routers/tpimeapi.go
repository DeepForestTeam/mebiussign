package routers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/timestamps"
)

func TimeApiGetCerrent(w http.ResponseWriter, r *http.Request) {
	ts := timestamps.TimeStampSignature{}
	err := ts.GetCurrent()
	if err != nil {
		log.Error("Can not create new time stamp!")
		fmt.Fprintf(w, `{"error":"Can not create time stamp", "ResultCode":500 }`)
		return
	} else {

	}
	time_stamp, _ := json.Marshal(ts)
	log.Debug("Get current time stamp")
	fmt.Fprintf(w, string(time_stamp))
}
