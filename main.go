package main

import (
	"flag"
	"github.com/DeepForestTeam/mobiussign/components"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/store"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/DeepForestTeam/mobiussign/restapi/routers"
)

func init() {
	log.Info("* Init main")
	config_file := flag.String("config", "conf/service.ini", "Define config file path.")
	config.GlobalConfig.LoadFromFile(*config_file)

}
func main() {
	log.Info("Starting MobiusSign™ ver.", components.APP_VERSION)
	err := store.GlobalStoreBarrel.ConnectDB()
	if err != nil {
		log.Fatal("BOLT?", err)
		panic(err)
	}
	log.Debug("Bolt/StromDB connected")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", routers.Index)
	http_port, err := config.GlobalConfig.GetString("HTTP_PORT")
	log.Fatal(http.ListenAndServe(":" + http_port, router))

}
