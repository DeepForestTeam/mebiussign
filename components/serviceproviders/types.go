package serviceproviders

import (
	"errors"
	"time"
)

const (
	ServicesStorage = "provider_store"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidPhoneNumber = errors.New("invalid phone format")
	ErrInvalidZipCode = errors.New("invalid zip code format")
	ErrInvalidCountry = errors.New("invalid country")
	ErrInvalidLogin = errors.New("invalid login")
	ErrInvalidApiKey = errors.New("invalid api key")
	ErrInvalidRsaKey = errors.New("invalid rsa public key format")
)

type ServiceProviderRow struct {
	Login             string
	Password          string
	ServiceId         string
	Contacts          struct {
				  Organisation       string
				  Person             struct {
							     Title    string
							     FullName string
							     Email    string
							     Phone    string
						     }
				  Country            string
				  ZipCode            string
				  StreetAddressLine1 string
				  StreetAddressLine2 string
				  City               string
			  }
	MetaInfo          struct {
				  RegiserDate time.Time
				  RegisterIp  string
			  }
	SignatureSettings struct {
				  UseStrong  bool
				  AllowedIps []string
			  }
}

type ServiceProviderApiKeys struct {
	ServiceId string
	ApiKey    string
	RsaPubKey string
}
