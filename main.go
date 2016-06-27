package main

import (
	"fmt"
	"github.com/DeepForestTeam/mebiussign/components/app"
	_ "github.com/DeepForestTeam/mebiussign/components/memstore"
)

func main() {
	fmt.Println("Starting MepiusSign(tm) ver.", app.APP_VERSION)
}
