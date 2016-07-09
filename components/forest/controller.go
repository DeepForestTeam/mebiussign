package forest

import (
	"net/http"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

func init() {
	log.Info("* Init Forest controller")
}

type Controll struct {
	Input    http.Request
	Output   http.ResponseWriter
	Url      string
	Domain   string
	Protocol string
	Method   string
	GetVars  map[string]string
	PostVars map[string]string
}

func (this *Controll)Process(r http.Request, w http.ResponseWriter) (err error) {
	return
}