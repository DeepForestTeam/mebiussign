package routers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/timestamps"
	"github.com/gorilla/mux"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
)

type TimeApiController struct {
	forest.Control
}

func init() {
	time_api_get := TimeApiController{}
	time_api_get.ThisName = "TimeApi™"
	forest.AddRouter("/api/time", &time_api_get)
	forest.AddRouterFunc("/api/time/{time_hash:[0-9ABCDEF]{64}}", TimeApiCheck)
}

func (this *TimeApiController)Get() {
	defer this.ServeJSON()
	ts := timestamps.TimeStampSignature{}
	err := ts.GetCurrent()
	if err != nil {
		log.Error("Can not create new time stamp!")
		fmt.Fprintf(this.Output, `{"error":"Can not create time stamp", "result_code":500 }`)
		return
	}

	this.Json = ts
}
func TimeApiCheck(w http.ResponseWriter, r *http.Request) {
	ts := timestamps.TimeStampSignature{}
	vars := mux.Vars(r)
	time_hash := vars["time_hash"]
	if time_hash == "" {
		log.Warning("No time hash present")
		log.Error("Can not read  time stamp!")
		fmt.Fprintf(w, `{ "error":"No time hash found", "result_code":404 }`)
		return
	}
	err := ts.GetBySign(time_hash)
	if err != nil {
		log.Error("Can not read  time stamp!")
		fmt.Fprintf(w, `{ "error":"No time hash found", "result_code":404 }`)
		return
	} else {

	}
	time_stamp, _ := json.Marshal(ts)
	log.Debug("Get time by stamp")
	fmt.Fprintf(w, string(time_stamp))
}
