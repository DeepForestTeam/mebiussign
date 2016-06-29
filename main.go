package main

import (
	"flag"
	"github.com/DeepForestTeam/mobiussign/components"
	_ "github.com/DeepForestTeam/mobiussign/components/memstore"
	_ "github.com/DeepForestTeam/mobiussign/components/timestamps"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"log"
	"github.com/DeepForestTeam/mobiussign/components/timestamps"
)

func init() {
	log.Println("* Init main")
	config_file := flag.String("config", "conf/service.ini", "Define config file path.")
	config.GlobalConfig.LoadFromFile(*config_file)
}
func main() {
	//Read config file
	log.Println("Starting MepiusSign(tm) ver.", components.APP_VERSION)
	ts := timestamps.TimeStampSignature{}
	log.Println(ts.GetCurrent())
}
