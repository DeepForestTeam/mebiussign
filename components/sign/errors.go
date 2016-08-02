package sign

import "errors"

var (
	ErrNoDataFound = errors.New("no data or data hash found")
	ErrEmptyDataBlock = errors.New("empty data block")
	ErrInvalidDataBlockFormat = errors.New("invalid data block format")
	ErrInvalidDataHash = errors.New("invalid data sha2 hash")
	ErrInvalidDataHashFormat = errors.New("data hash not SHA512")
	ErrInvalidDataFormat = errors.New("invalid data format")

	ErrSignNotFound = errors.New("signature not found")
)
