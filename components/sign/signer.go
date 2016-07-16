package sign

import (
	"fmt"
	"sync"
	"regexp"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"encoding/base64"
	"github.com/DeepForestTeam/mobiussign/components/timestamps"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/DeepForestTeam/mobiussign/components/store"
)

type MobiusSigner struct {
	SignRequest  SignatureRequest
	SignRow      SignatureRow
	SignResponse SignatureResponse
	mux          sync.Mutex
}

func (this *MobiusSigner)ProcessQuery() (err error) {
	err = this.processData()
	if err != nil {
		return
	}
	return
}
func (this *MobiusSigner)processData() (err error) {
	this.SignRow.RsaSignature = "n/a"
	if this.SignRequest.DataBlock == "" && this.SignRequest.DataHash == "" {
		return ErrNoDataFound
	} else if this.SignRequest.DataBlock != "" {
		err = this.prepareData()
		if err != nil {
			return
		}
		if this.SignRequest.DataHash == "" {
			err = this.createRequestDataHash()
		} else {
			err = this.checkRequestDataHash()
		}
		if err != nil {
			return
		}
	}
	this.SignRow.ServiceId = this.SignRequest.ServiceId
	this.SignRow.ObjectId = this.SignRequest.ObjectId
	this.SignRow.ConsumerId = this.SignRequest.ConsumerId

	this.SignRow.DataUrl = this.SignRequest.DataUrl
	this.SignRow.DataNote = this.SignRequest.DataNote
	this.SignRow.DataHash = this.SignRequest.DataHash
	this.SignRow.DataBlock = this.SignRequest.DataBlock

	mobius_time := timestamps.TimeStampSignature{}
	err = mobius_time.GetCurrent()
	if err != nil {
		return
	}
	this.SignRow.TimeStamp = mobius_time.TimeStamp
	this.SignRow.UnixTimeStamp = mobius_time.UnixTimeStamp
	this.SignRow.TimeStampHash = mobius_time.MobiusTime

	this.SignRow.PepperHash = fmt.Sprintf("%X", generatePepper())
	//reserved
	last_record := SignatureRow{}
	last_key, err := store.Last(MobiusStorage, &last_record)
	base_salt, err := config.GetString("BASE_WORLD_HASH")
	if err != nil {
		return err
	}
	if err != nil {
		if err != store.ErrKeyNotFound && err != store.ErrSectionNotFound {
			log.Error("Storage error:", err)
			return
		}
		this.SignRow.SaltId = ""
		this.SignRow.SaltHash = base_salt
	} else {
		this.SignRow.SaltId = last_key
		this.SignRow.SaltHash = last_record.MobiusSignature
	}
	err = this.Sign()
	if err != nil {
		log.Error("Can create MobiusSign:", err)
	}
	key := this.SignRow.MobiusSignature
	err = store.Set(MobiusStorage, key, &this.SignRow)
	return
}
func (this *MobiusSigner)prepareData() (err error) {
	err = nil
	if this.SignRequest.DataBlockFormat != DataBlockBase64 && this.SignRequest.DataBlockFormat != DataBlockString &&this.SignRequest.DataBlockFormat != DataBlockHex && this.SignRequest.DataBlockFormat != "" {
		log.Error("Error:", ErrInvalidDataBlockFormat)
		return ErrInvalidDataBlockFormat
	}
	if this.SignRequest.DataBlockFormat == "" {
		this.SignRequest.DataBlockFormat = DataBlockString
	}
	if this.SignRequest.DataHash != "" && len(this.SignRequest.DataHash) != 64 {
		return ErrInvalidDataHashFormat
	}
	return
}

func (this *MobiusSigner)Sign() (err error) {
	salt_hash, err := decodeHex(this.SignRow.SaltHash)
	if err != nil {
		log.Error("Can not decode salt hash:", err)
		return
	}
	pepper, err := decodeHex(this.SignRow.PepperHash)
	if err != nil {
		log.Error("Can not decode pepper hash:", err)
		return
	}
	service_id := []byte(this.SignRow.ServiceId)
	object_id := []byte(this.SignRow.ObjectId)
	consumer_id := []byte(this.SignRow.ConsumerId)
	data_hash, err := decodeHex(this.SignRow.DataHash)
	if err != nil {
		log.Error("Can not decode data hash:", err)
		return
	}
	mobius_time, err := decodeHex(this.SignRow.TimeStampHash)
	if err != nil {
		log.Error("Can not decode time hash:", err)
		return
	}
	signin_block := joinByteBlocks(salt_hash, pepper, service_id, object_id, consumer_id, data_hash, mobius_time)
	log.Debug("Sign block len:", len(signin_block))
	this.SignRow.MobiusSignature = calculateHash(signin_block)
	return
}

func (this *MobiusSigner)createRequestDataHash() (err error) {
	data_block, err := this.decodeRequestDataBlock()
	if err != nil {
		return
	}
	data_hash := calculateHash(data_block)
	this.SignRequest.DataHash = data_hash
	return
}

func (this *MobiusSigner)checkRequestDataHash() (err error) {
	data_block, err := this.decodeRequestDataBlock()
	if err != nil {
		return
	}
	calculated_data_hash := calculateHash(data_block)
	if this.SignRequest.DataHash != calculated_data_hash {
		return ErrInvalidDataHash
	}
	return
}
func (this *MobiusSigner)decodeRequestDataBlock() (data_block []byte, err error) {
	switch this.SignRequest.DataBlockFormat {
	case DataBlockBase64:
		data_block, err = base64.StdEncoding.DecodeString(this.SignRequest.DataBlock)
	case DataBlockString:
		data_block = []byte(this.SignRequest.DataBlock)
	case DataBlockHex:
		data_block, err = decodeHex(this.SignRequest.DataBlock)
	default:
		return data_block, ErrInvalidDataBlockFormat
	}
	return
}

func calculateHash(bytes_block []byte) (hash string) {
	hasher := sha512.New()
	hasher.Write(bytes_block)
	bin_hash := hasher.Sum(nil)
	hash = fmt.Sprintf("%X", bin_hash)
	return
}
func decodeHex(input string) (bytes_block []byte, err error) {
	clear_whitespaces := regexp.MustCompile(`\s`)
	final := clear_whitespaces.ReplaceAllString(input, "")
	bytes_block, err = hex.DecodeString(final)
	return
}
func generatePepper() (pepper []byte) {
	pepper = make([]byte, 8)
	rand.Read(pepper)
	return
}
func joinByteBlocks(b... []byte) (joined []byte) {
	for _, block := range b {
		joined = append(joined, block...)
	}
	return
}