package forest

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

var router = mux.NewRouter().StrictSlash(false)

func init() {
	log.Info("* Init Forest router")
}

type Controller interface {
	Process() error
}

func AddRouterFunc(uri string, f func(http.ResponseWriter, *http.Request)) {
	log.Debug("Register flat handler for:", uri)
	router.HandleFunc(uri, f)
}
func AddRouterObject(uri string, handler Control) {
	router.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		handler.Input = *r
		handler.Output = w
		handler.Process()
	})
}