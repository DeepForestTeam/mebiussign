package routers

import (
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/restapi/controls"
)

func init() {
	sign_api := controls.SignController{}
	sign_api.ThisName = "SignApi™"
	forest.AddRouter("/api/sign", &sign_api)
	forest.AddRouter("/api/sign/{sign_hash:[0-9A-Fa-f]{16}}", &sign_api)
	forest.AddRouter("/api/sign/{sign_hash:[0-9A-Fa-f]{128}}", &sign_api)
}

