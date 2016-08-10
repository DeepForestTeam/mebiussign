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
	"errors"
	"html"
	"strings"
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
		error_msg := err.Error()
		this.Data = ErrorMessage{
			Result:error_msg,
			Note:error_msg,
			ResultCode:406,
		}
		return
	}
	this.prepareSignRequest(&signer.SignRequest)
	err = signer.ProcessQuery()
	if err != nil {
		log.Error("Invalid request data:", err)
		error_msg := err.Error()
		this.Data = ErrorMessage{
			Result:error_msg,
			Note:error_msg,
			ResultCode:406,
		}
		return
	}
	signer.SignResponse.Result = "OK"
	this.Data = signer.SignResponse
}

func (this *SignController)validateSignRequest(sign_request *sign.SignatureRequest) (err error) {
	if sign_request.DataHash == "" && sign_request.DataBlock == "" {
		return errors.New("no data for sign")
	} else if sign_request.DataHash != "" {
		if !validation.IsSha512(sign_request.DataHash) {
			return errors.New("invalid data hash")
		}
	}
	if sign_request.ServiceId != "" {
		if !validation.IsAlphaNumeric(sign_request.ServiceId) {
			return errors.New("invalid chars in service_id")
		}
	}
	if sign_request.ObjectId != "" {
		if !validation.IsAlphaNumeric(sign_request.ObjectId) {
			return errors.New("invalid chars in object_id")
		}
	}
	if sign_request.ConsumerId != "" {
		if !validation.IsAlphaNumeric(sign_request.ConsumerId) {
			return errors.New("invalid chars in consumer_id")
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
		if sign_request.DataBlockFormat == sign.DataBlockBase64 {
			if !validation.IsBase64(sign_request.DataBlock) {
				return errors.New("invalid data block: not base64")
			}
		}
		if sign_request.DataBlockFormat == sign.DataBlockHex {
			if !validation.IsHex(sign_request.DataBlock) {
				return errors.New("invalid data block: not valid hex")
			}
		}
	}
	return
}

func (this *SignController)prepareSignRequest(sign_request *sign.SignatureRequest) {
	if sign_request.DataNote != "" {
		sign_request.DataNote = html.EscapeString(sign_request.DataNote)
	}
	if sign_request.DataHash != "" {
		sign_request.DataHash = strings.ToUpper(sign_request.DataHash)
	}
	if sign_request.ServiceSign == "" {
		sign_request.ServiceId = ""
	}
}