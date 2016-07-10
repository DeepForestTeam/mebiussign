package routers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/timestamps"
	"github.com/gorilla/mux"
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
func TimeApiCheck(w http.ResponseWriter, r *http.Request) {
	ts := timestamps.TimeStampSignature{}
	vars := mux.Vars(r)
	time_hash := vars["time_hash"]
	if time_hash == "" {
		log.Warning("No time hash present")
		log.Error("Can not create new time stamp!")
		fmt.Fprintf(w, `{ "error":"No time hash present", "ResultCode":404 }`)
		return
	}
	err := ts.GetBySign(time_hash)
	if err != nil {
		log.Error("Can not read new time stamp!")
		fmt.Fprintf(w, `{ "error":"No time hash present", "ResultCode":404 }`)
		return
	} else {

	}
	time_stamp, _ := json.Marshal(ts)
	log.Debug("Get current time stamp")
	fmt.Fprintf(w, string(time_stamp))
}
