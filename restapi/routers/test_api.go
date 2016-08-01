package routers

import (
	"net/http"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/sign"
	"github.com/DeepForestTeam/mobiussign/components/serviceproviders"
)

func InternalTest(w http.ResponseWriter, r *http.Request) {
	var test_answer map[string]interface{}
	test_answer = make(map[string]interface{})
	test_object := sign.MobiusSigner{}
	test_object.SignRequest.DataBlock = "The Ultimate Question of Life, the Universe, and Everything:42"
	test_object.SignRequest.DataNote = "The Ultimate Question of Life, the Universe, and Everything"
	test_object.SignRequest.DataUrl = "https://en.wikipedia.org/wiki/Phrases_from_The_Hitchhiker%27s_Guide_to_the_Galaxy#Answer_to_the_Ultimate_Question_of_Life.2C_the_Universe.2C_and_Everything_.2842.29"
	err := test_object.ProcessQuery()
	if err != nil {
		log.Error(err)
	}
	test_answer["sing"] = test_object
	json_string, _ := json.MarshalIndent(&test_answer, "", "  ")
	w.Write(json_string)
}

func InternalTest2(w http.ResponseWriter, r *http.Request) {
	var test_answer map[string]interface{}
	test_answer = make(map[string]interface{})
	object := serviceproviders.ServiceProviderRow{}

	object.ServiceId = "DEEDSCHAIN01"
	object.Contacts.Country = "UA"
	object.Contacts.Organisation = "DeedsChain Ink"
	object.Contacts.Person.FullName = "Ashot Ishimbaevich"
	object.Contacts.Person.Email = "team@deedschain.com"
	object.Contacts.Person.Phone = "+102545997783"
	object.Contacts.ZipCode = "12300"
	object.Contacts.City = "Kiyv"

	test_answer["SERVICE_RECORD"] = object
	json_string, _ := json.MarshalIndent(&test_answer, "", "  ")
	w.Write(json_string)
}

func init() {
	log.Warning("* Init Test API")
	forest.AddRouterFunc("/api/test", InternalTest)
	forest.AddRouterFunc("/api/test/1", InternalTest)
	forest.AddRouterFunc("/api/test/2", InternalTest2)
}