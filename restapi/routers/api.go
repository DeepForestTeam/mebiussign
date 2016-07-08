package routers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/components"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MobiusSign™ API")
}
func Ping(w http.ResponseWriter, r *http.Request) {
	var ping_answer map[string]interface{}
	ping_answer = make(map[string]interface{})
	ping_answer["ping"] = "pong"
	ping_answer["version"] = components.APP_VERSION
	ping_answer["service"] = "MobiusSign™ API"
	ping_string, _ := json.Marshal(ping_answer)
	fmt.Fprintf(w, string(ping_string))
}
