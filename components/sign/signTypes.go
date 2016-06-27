package sign

import "time"

type SignatureRow struct {
	RowIndex        string        `json:"key"`
	TimeStamp       time.Time     `json:"time_stamp"`
	TimeStampHash   string        `json:"time_stamp_hash"`
	//USER PROVIDED DATA
	DataUrl         string        `json:"data_url"`
	DataNote        string        `json:"data_note"`
	DataSignature   string        `json:"data_signature"`
	DataBlock       []byte        `json:"data_block"`
	//System given info
	PrevKey         string        `json:"prev_key"`
	PrevSignature   string        `json:"prev_signature"`
	MobiusSignature string        `json:"mobius_sign"`
}

type MobiusTape struct {
	LastKey    string           `json:"last_key"`
	LastIndex  string           `json:"last_index"`
	KeyIndex   map[string]int64 `json:"key_index"`
	MobiusRows []SignatureRow   `json:"mobius_rows"`
}

