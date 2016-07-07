package config

import (
	"os"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"bufio"
	"strings"
	"strconv"
)

type Config struct {
	store map[string]string
	cache map[string]interface{}
}

var GlobalConfig Config

func init() {
	log.Debug("* Init config")
	GlobalConfig.store = make(map[string]string)
}

// Load key-value pairs from ini-like config file
// key=value
// Any line, starting # - comment

func (this *Config)LoadFromFile(config_file string) (err error) {
	err = nil
	file, err := os.Open(config_file)
	if err != nil {
		log.Fatal("Can not open config file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") && strings.Index(line, "=") != -1 {
			line_parts := strings.Split(line, "=")
			this.store[strings.TrimSpace(line_parts[0])] = strings.TrimSpace(line_parts[1])
		}
	}
	return
}

type ConfigError struct {
	ErrorMesssage string
	ErrorInKey    string
}

func (this ConfigError)Error() string {
	return this.ErrorMesssage
}
//Get string value from config
func (this *Config)GetString(key string) (value string, err error) {
	err = nil
	value, ok := this.store[key]
	if !ok {
		err := ConfigError{ErrorMesssage:"Key not found:" + key, ErrorInKey:key}
		return "", err
	}
	value = strings.Trim(value, `"`)
	return
}
//Get int64 value from config
func (this *Config)GetInt64(key string) (i int64, err error) {
	err = nil
	value, ok := this.store[key]
	if !ok {
		err := ConfigError{ErrorMesssage:"Key not found:" + key, ErrorInKey:key}
		return 0, err
	}
	i, err = strconv.ParseInt(value, 10, 64)
	return
}
//Get int32 value from config
func (this *Config)GetInt32(key string) (i int32, err error) {
	err = nil
	value, ok := this.store[key]
	if !ok {
		err := ConfigError{ErrorMesssage:"Key not found:" + key, ErrorInKey:key}
		return 0, err
	}
	i64, err := strconv.ParseInt(value, 10, 32)
	i = int32(i64)
	return i, err
}
//Get boolean value from config
func (this *Config)GetBool(key string) (b bool, err error) {
	err = nil
	value, ok := this.store[key]
	if !ok {
		err := ConfigError{ErrorMesssage:"Key not found:" + key, ErrorInKey:key}
		return false, err
	}
	b, err = strconv.ParseBool(value)
	return
}
//Get float value from config
func (this *Config)GetFloat(key string) (f float64, err error) {
	err = nil
	value, ok := this.store[key]
	if !ok {
		err := ConfigError{ErrorMesssage:"Key not found:" + key, ErrorInKey:key}
		return 0, err
	}
	f, err = strconv.ParseFloat(value, 64)
	return
}
