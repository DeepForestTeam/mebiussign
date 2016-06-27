package main

import (
	"fmt"
	"github.com/DeepForestTeam/mobiussign/components/app"
	_ "github.com/DeepForestTeam/mobiussign/components/memstore"
)

func main() {
	fmt.Println("Starting MepiusSign(tm) ver.", app.APP_VERSION)
}
