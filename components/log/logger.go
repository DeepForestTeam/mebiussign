package log

import (
	"github.com/DeepForestTeam/go-logging"
	"os"
)

var log = logging.MustGetLogger("example")

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05} %{shortfile}` + "\t" + `: %{callpath}` + "\t"+ `| %{level:.4s}` + "\t| " + `%{message}%{color:reset}`,
	//`%{color}%{time:15:04:05.000} %{shortfunc} ` + "\t" + `▶ %{level:.4s} %{id:03x} %{message}%{color:reset}`,
)
var backend *logging.LogBackend

func init() {
	backend = logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	log.ExtraCalldepth=1
	log.Info("* Init loger")
}

func Debug(v...interface{}) {
	log.Debug(v...)
}
func Info(v...interface{}) {
	log.Info(v...)
}
func Fatal(v...interface{}) {
	log.Fatal(v...)
}
func Panic(v...interface{}) {
	log.Panic(v...)
}
func Critical(v...interface{}) {
	log.Critical(v...)
}
func Error(v...interface{}) {
	log.Error(v...)
}
func Warning(v...interface{}) {
	log.Warning(v...)
}
func Notice(v...interface{}) {
	log.Notice(v...)
}