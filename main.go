package main

import (
	"fmt"
	"flag"
	"github.com/DeepForestTeam/mobiussign/components"
	_ "github.com/DeepForestTeam/mobiussign/components/memstore"
	"github.com/DeepForestTeam/mobiussign/components/config"
)

func main() {
	//Read config file
	config_file := flag.String("config", "conf/service.ini", "Define config file path.")
	config.GlobalConfig.LoadFromFile(*config_file)
	fmt.Println("Starting MepiusSign(tm) ver.", components.APP_VERSION)
}
