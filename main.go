package main

import (
	"log"
	"flag"
	"github.com/DeepForestTeam/mobiussign/components"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/timestamps"
	_ "github.com/DeepForestTeam/mobiussign/components/memstore"
)

func init() {
	log.Println("* Init main")
	config_file := flag.String("config", "conf/service.ini", "Define config file path.")
	config.GlobalConfig.LoadFromFile(*config_file)
}
func main() {
	log.Println("Starting MepiusSign(tm) ver.", components.APP_VERSION)
	ts := timestamps.TimeStampSignature{}
	log.Println(ts.GetCurrent())
}
