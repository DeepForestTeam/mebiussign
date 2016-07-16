package routers

import (
	"net/http"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/sign"
)

func InternalTest(w http.ResponseWriter, r *http.Request) {
	var test_answer map[string]interface{}
	test_answer = make(map[string]interface{})
	test_object := sign.MobiusSigner{}
	test_object.SignRequest.DataBlock = "The Ultimate Question of Life, the Universe, and Everything:42"
	test_object.SignRequest.DataNote="The Ultimate Question of Life, the Universe, and Everything"
	test_object.SignRequest.DataUrl="https://en.wikipedia.org/wiki/Phrases_from_The_Hitchhiker%27s_Guide_to_the_Galaxy#Answer_to_the_Ultimate_Question_of_Life.2C_the_Universe.2C_and_Everything_.2842.29"
	err := test_object.ProcessQuery()
	if err != nil {
		log.Error(err)
	}
	test_answer["sing"] = test_object
	json_string, _ := json.MarshalIndent(&test_answer, "", "  ")
	w.Write(json_string)
}

func init() {
	log.Warning("* Init Test API")
	forest.AddRouterFunc("/api/test", InternalTest)
}