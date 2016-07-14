package sign

import "time"

type SignatureRow struct {
	//
	SignId          string        `json:"sign_id"`
	RowId           uint64        `json:"row_id"`
	BlockId         string        `json:"block_id"`
	//USER PROVIDED DATA
	DataUrl         string        `json:"data_url"`
	DataNote        string        `json:"data_note"`
	DataHash        string        `json:"data_hash"`
	DataBlock       []byte        `json:"data_block"`
	//SYSTEM GIVEN INFO
	TimeStamp       time.Time     `json:"time_stamp"`
	TimeStampHash   string        `json:"time_hash"`
	SaltKey         string        `json:"salt_key"`
	SaltHash        string        `json:"salt_hash"`
	MobiusSignature string        `json:"mobius_sign"`
}

type SignatureRequest struct {
	//Service and object identification
	ServiceID     string        `json:"service_id"`
	ObjectID      string        `json:"object_id"`
	ConsumerID    string        `json:"object_id"`
	//DATA SECTION
	DataUrl       string        `json:"data_url"`
	DataNote      string        `json:"data_note"`
	DataSignature string        `json:"data_signature"`
	DataBlock     []byte        `json:"data_block"`
	//SYSTEM GIVEN INFO
	ServiceSign   string        `json:"service_sign"`
}

