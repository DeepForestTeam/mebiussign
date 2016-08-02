package routers

import (
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/restapi/controls"
)

func init() {
	time_api_get := controls.TimeApiController{}
	time_api_get.ThisName = "TimeApi™"
	forest.AddRouter("/api/time", &time_api_get)
	forest.AddRouter("/api/time/{time_hash:[0-9A-F]{64}}", &time_api_get)
}

