package controls

import (
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/DeepForestTeam/mobiussign/components/sign"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/validation"
	"io/ioutil"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"fmt"
	"encoding/json"
	"github.com/iris-contrib/errors"
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
	mobius_sign.SignResponse.Result = "OK"
	this.Data = mobius_sign.SignResponse
}
//Create signature
func (this *SignController)Post() {
	defer this.ServeJSON()
	this.Output.Header().Set("Access-Control-Allow-Origin", "*")
	signer := sign.MobiusSigner{}
	body, err := ioutil.ReadAll(this.Input.Body)
	max_len, err := config.GetInt64("SIGN_MAX_REQUEST")
	if err != nil {
		max_len = 32768
		log.Warning("Can not read, default value used(32768)")
	}
	if len(body) > int(max_len) {
		log.Error("Content too long:", len(body))
		this.Data = ErrorMessage{
			Result:"Content too long.",
			Note: fmt.Sprintf("Max sign request body size %d bytes", max_len),
			ResultCode:406,
		}
		return
	}
	err = json.Unmarshal(body, &signer.SignRequest)
	if err != nil {
		log.Error("Invalid JSON:", len(body))
		this.Data = ErrorMessage{
			Result:"Invalid JSON",
			ResultCode:406,
		}
		return
	}
	err = this.validateSignRequest(&signer.SignRequest)
	if err != nil {
		log.Error("Invalid request data:", err)
		this.Data = ErrorMessage{
			Result:err.Error(),
			ResultCode:406,
		}
		return
	}
	this.Data = signer.SignResponse
}

func (this *SignController)validateSignRequest(sign_request *sign.SignatureRequest) (err error) {
	if sign_request.DataHash == "" && sign_request.DataHash == "" {
		return errors.New("no data for sign")
	} else if sign_request.DataHash != "" {
		if !validation.IsSha512(sign_request.DataHash) {
			return errors.New("invalid data hash")
		}
	}
	if sign_request.DataUrl != "" {
		if !validation.IsUrl(sign_request.DataUrl) {
			return errors.New("invalid data url")
		}
	}
	if sign_request.DataBlockFormat != "" && sign_request.DataBlock != "" {
		if sign_request.DataBlockFormat != sign.DataBlockBase64 && sign_request.DataBlockFormat != sign.DataBlockHex && sign_request.DataBlockFormat != sign.DataBlockString {
			return errors.New("unknown data flock format")
		}
		if sign_request.DataBlockFormat == sign.DataBlockBase64{

		}
	}
	return
}