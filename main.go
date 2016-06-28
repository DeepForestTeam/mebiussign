package main

import (
	"fmt"
	"github.com/DeepForestTeam/mobiussign/components"
	_ "github.com/DeepForestTeam/mobiussign/components/memstore"
)

func main() {
	fmt.Println("Starting MepiusSign(tm) ver.", components.APP_VERSION)
}
