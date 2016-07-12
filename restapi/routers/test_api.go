package routers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mebiussign/components/sign"
)

func InternalTest(w http.ResponseWriter, r *http.Request) {
	var test_answer map[string]interface{}
	test_answer = make(map[string]interface{})
	test_answer["test"] = "BoltTest:Last()"
	test_object := sign.SignatureRow{}
	test_object.DataBlock=[]byte("8399EB32571640591483C62C0F223BAAE87F658E2C20E2EF480E4B10EC25C396")
	test_answer["sing"] = test_object
	json_string, _ := json.Marshal(&test_answer)
	fmt.Fprintf(w, string(json_string))
}

func init() {
	log.Warning("* Init Test API")
	forest.AddRouterFunc("/api/test", InternalTest)
}