package sign

import "time"

type SignatureRow struct {
	//
	RowIndex        string    `json:"key"`
	//USER PROVIDED DATA
	DataUrl         string `json:"data_url"`
	DataNote        string `json:"data_note"`
	DataHash        string `json:"data_hash"`
	DataBlock       []byte `json:"data_block"`
	//System given info
	TimeStamp       time.Time `json:"time_stamp"`
	TimeStampHash   string    `json:"time_hash"`
	PrevKey         string `json:"prev_key"`
	PrevSignature   string `json:"prev_signature"`
	MobiusSignature string `json:"mobius_sign"`
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
	//Security Section
	ServiceSign   string        `json:"service_sign"`
}

type MobiusTape struct {
	LastKey    string           `json:"last_key"`
	LastIndex  string           `json:"last_index"`
	KeyIndex   map[string]int64 `json:"key_index"`
	MobiusRows []SignatureRow   `json:"mobius_rows"`
}
