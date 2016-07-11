package main

import (
	"github.com/DeepForestTeam/mobiussign/components"
	"github.com/DeepForestTeam/mobiussign/components/log"
	_ "github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/store"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"

	_ "github.com/DeepForestTeam/mobiussign/restapi/routers"
)

func init() {
	log.Info("* Init main")
}
func main() {
	log.Info("Starting MobiusSign™ ver.", components.APP_VERSION)
	err := store.ConnectDB()
	if err != nil {
		log.Fatal("BOLT?", err)
		panic(err)
	}
	log.Debug("Bolt/StromDB connected")
	err = forest.StartServer()
	log.Fatal(err)
}
