package forest

import (
	"net/http"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

func init() {
	log.Info("* Init Forest controller")
}

type Control struct {
	Input  http.Request
	Output http.ResponseWriter
	Ctx    Context
	DoPost func()
	DoGet  func()
}
type Context struct {
	Url      string
	Domain   string
	Protocol string
	Method   string
	UrlVars  map[string]string
	GetVars  map[string]string
	PostVars map[string]string
}

func (this *Control)Process() (err error) {
	return
}

func (this *Control)Get() {

}
func (this *Control)Post() {

}
func (this *Control)Put() {

}
func (this *Control)Delete() {

}