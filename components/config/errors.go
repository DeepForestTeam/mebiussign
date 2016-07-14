package config

import "errors"

var (
	ErrorKeyNotFound = errors.New("config value not found")
	ErrorIncorrectValueFormat = errors.New("config value not found")
)
