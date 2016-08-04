package sign

const (
	MobiusStorage = "sign_store"

	DataBlockBase64 = "base64"
	DataBlockString = "string"
	DataBlockHex = "hex"
)

type SignatureRow struct {
	SignId          string        `json:"sign_id"`
	RowId           int64         `json:"row_id"`
	//Reserved for CloudMobius
	BlockId         string        `json:"block_id"`
	//USER PROVIDED DATA
	ServiceId       string        `json:"service_id"`
	ObjectId        string        `json:"object_id"`
	ConsumerId      string        `json:"consumer_id"`
	DataUrl         string        `json:"data_url"`
	DataNote        string        `json:"data_note"`
	DataHash        string        `json:"data_hash"`
	DataBlock       string        `json:"data_block"`
	DataBlockFormat string        `json:"data_format"`
	//SYSTEM GIVEN INFO
	TimeStamp       string        `json:"time"`
	UnixTimeStamp   int64         `json:"unix_time"`
	TimeStampHash   string        `json:"time_hash"`
	//Sign info
	SaltId          string        `json:"salt_id"`
	SaltHash        string        `json:"salt_hash"`
	PepperHash      string        `json:"pepper_hash"`
	MobiusSignature string        `json:"mobius_sign"`
	RsaSignature    string        `json:"rsa_sign"`
}

type SignatureRequest struct {
	//Service and object identification
	//ServiceIdmust be 0 for public or personal use
	ServiceId       string        `json:"service_id"`
	ObjectId        string        `json:"object_id"`
	ConsumerId      string        `json:"consumer_id"`
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
	Result          string        `json:"result"`
	SignId          string        `json:"sign_id"`
	RowId           int64         `json:"row_id"`
	//Reserved for CloudMobius
	BlockId         string        `json:"block_id"`

	ServiceId       string        `json:"service_id"`
	ObjectId        string        `json:"object_id"`
	ConsumerId      string        `json:"consumer_id"`

	DataUrl         string        `json:"data_url"`
	DataNote        string        `json:"data_note"`
	DataHash        string        `json:"data_hash"`
	DataBlock       string        `json:"data_block"`
	DataBlockFormat string        `json:"data_format"`

	TimeStamp       string        `json:"time"`
	UnixTimeStamp   int64         `json:"unix_time"`
	TimeStampHash   string        `json:"time_hash"`

	SaltId          string        `json:"salt_id"`
	SaltHash        string        `json:"salt_hash"`
	PepperHash      string        `json:"pepper_hash"`
	MobiusSignature string        `json:"mobius_sign"`
	RsaSignature    string        `json:"rsa_sign"`
}