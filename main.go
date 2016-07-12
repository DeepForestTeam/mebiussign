package main

import (
	"github.com/DeepForestTeam/mobiussign/components"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	_ "github.com/DeepForestTeam/mobiussign/components/config"
	_ "github.com/DeepForestTeam/mobiussign/restapi/routers"
	"github.com/DeepForestTeam/mobiussign/components/store"
)

func init() {
	log.Info("* Init main")
}
func main() {
	log.Info("Starting MobiusSign™ ver.", components.APP_VERSION)
	err := store.ConnectDB()
	if err != nil {
		log.Fatal("Can not connect to storage:", err)
		panic(err)
	}
	log.Debug("Bolt/StromDB connected")
	err = forest.StartServer()
	log.Fatal(err)
}
