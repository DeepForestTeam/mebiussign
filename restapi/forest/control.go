// Warper over Gorilla Mux package
package forest

import (
	"net/http"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"encoding/xml"
)

func init() {
	log.Info("* Init Forest controller")
}

type Controller interface {
	Process(http.ResponseWriter, *http.Request)
	Get()
	Post()
	RenderTemplate()
}

type Control struct {
	ThisName      string
	Input         *http.Request
	Output        http.ResponseWriter
	Context       Context
	//Custom handlers
	PostHandler   func(*Control)
	GetHandler    func(*Control)
	PutHandler    func(*Control)
	DeleteHandler func(*Control)
	OptionHandler func(*Control)
	//Output settings
	AutoRender    bool
	Template      string
	Json          interface{}
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

func (this *Control)Process(w http.ResponseWriter, r *http.Request) {
	this.Output = w
	this.Input = r
	this.Context.UrlVars = mux.Vars(r)
	return
}

func (this *Control)Get() {
	if this.GetHandler != nil {
		this.GetHandler(this)
	} else {
		log.Warning("Call default Get method", this.ThisName)
		this.Output.WriteHeader(http.StatusMethodNotAllowed)
		this.Output.Write([]byte("405 Method Not Allowed"))
	}
}
func (this *Control)Post() {
	if this.PostHandler != nil {
		this.PostHandler(this)
	} else {
		log.Warning("Call default Post method", this.ThisName)
		this.Output.WriteHeader(http.StatusMethodNotAllowed)
		this.Output.Write([]byte("405 Method Not Allowed"))
	}
}
func (this *Control)Put() {
	this.Output.WriteHeader(http.StatusMethodNotAllowed)
	this.Output.Write([]byte("405 Method Not Allowed"))
}
func (this *Control)Delete() {
	this.Output.WriteHeader(http.StatusMethodNotAllowed)
	this.Output.Write([]byte("405 Method Not Allowed"))
}
func (this *Control)Option() {
	this.Output.WriteHeader(http.StatusMethodNotAllowed)
	this.Output.Write([]byte("405 Method Not Allowed"))
}
func (this *Control)Error(http_error_code string) {

}

func (this *Control)parseRequestData() (err error) {
	err = nil
	return
}

//Support mehtods
func (this *Control)ServeJSON() {
	this.AutoRender = false
	json_string, _ := json.MarshalIndent(this.Json, "", "  ")
	this.Output.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(this.Output, string(json_string))
	return
}
func (this *Control)ServeXML() {
	this.AutoRender = false
	xml_string, _ := xml.MarshalIndent(this.Json, "", "  ")
	this.Output.Header().Set("Content-Type", "application/xml; charset=utf-8")
	fmt.Fprintf(this.Output, string(xml_string))
	return
}
func (this *Control)RenderTemplate() {
	if this.AutoRender {
		//@todo template renderer
	}
	return
}