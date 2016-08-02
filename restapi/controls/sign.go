package controls

import (
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/components/sign"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

type SignController struct {
	forest.Control
}

func (this *SignController)Get() {
	defer this.ServeJSON()
	sign_hash := this.Context.UrlVars["sign_hash"]
	if sign_hash == "" {
		this.Data = ErrorMessage{
			Result:"Method not allowed",
			ResultCode:405,
		}
		return
	}
	mobius_sign := sign.MobiusSigner{}
	err := mobius_sign.Check(sign_hash)
	if err != nil {
		log.Error("Error, while sign check:", err)
		this.Data = ErrorMessage{
			Result:"Signature not found",
			ResultCode:404,
		}
		return
	}
	this.Data = mobius_sign.SignResponse
}
func (this *SignController)Post() {
	defer this.ServeJSON()
	sign_result := sign.MobiusSigner{}
	this.Data = sign_result
}