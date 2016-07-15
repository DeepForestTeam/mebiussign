package sign

import "time"

const (
	SignatureStorage = "sign_store"
	SignatereIndex = "sign_index"
)

type SignatureRow struct {
	SignId          string        `json:"sign_id"`
	RowId           uint64        `json:"row_id"`
	//Reserved for CloudMobius
	BlockId         string        `json:"block_id"`
	//USER PROVIDED DATA
	ServiceId       string        `json:"service_id"`
	ObjectID        string        `json:"object_id"`
	ConsumerId      string        `json:"consumer_id"`
	DataUrl         string        `json:"data_url"`
	DataNote        string        `json:"data_note"`
	DataHash        string        `json:"data_hash"`
	DataBlock       []byte        `json:"data_block"`
	//SYSTEM GIVEN INFO
	TimeStamp       time.Time     `json:"time_stamp"`
	TimeStampHash   string        `json:"time_hash"`
	SaltKey         string        `json:"salt_key"`
	SaltHash        string        `json:"salt_hash"`
	PepperHash      string        `json:"pepper_hash"`
	MobiusSignature string        `json:"mobius_sign"`
}

type SignatureRequest struct {
	//Service and object identification
	//ServiceId must be 0 for public or personal use
	ServiceID       string        `json:"service_id"`
	ObjectID        string        `json:"object_id"`
	ConsumerID      string        `json:"consumer_id"`
	//DATA SECTION
	DataUrl         string        `json:"data_url"`
	DataNote        string        `json:"data_note"`
	DataHash        string        `json:"data_hash"`
	DataBlock       string        `json:"data_block"`
	DataBlockFormat string        `json:"data_format"`
	//SYSTEM GIVEN INFO
	ServiceSign     string        `json:"service_sign"`
}

type SignatureResponse struct {

}