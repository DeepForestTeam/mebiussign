package config

import (
	"os"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"bufio"
	"strings"
	"strconv"
	"flag"
)

type Config struct {
	store map[string]string
	cache map[string]interface{}
}

var GlobalConfig Config

func init() {
	log.Info("* Init config")
	config_file := flag.String("config", "conf/service.ini", "Define config file path.")
	GlobalConfig.store = make(map[string]string)
	GlobalConfig.LoadFromFile(*config_file)
}

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

//Get string value from config
func (this *Config)GetString(key string) (value string, err error) {
	err = nil
	value, ok := this.store[key]
	if !ok {
		err := ErrorKeyNotFound
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
		err := ErrorKeyNotFound
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
		err := ErrorKeyNotFound
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
		err := ErrorKeyNotFound
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
		err := ErrorKeyNotFound
		return 0, err
	}
	f, err = strconv.ParseFloat(value, 64)
	return
}

func GetString(key string) (value string, err error) {
	return GlobalConfig.GetString(key)
}

func GetInt64(key string) (i int64, err error) {
	return GlobalConfig.GetInt64(key)
}

func GetInt32(key string) (i int32, err error) {
	return GlobalConfig.GetInt32(key)
}

func GetBool(key string) (b bool, err error) {
	return GlobalConfig.GetBool(key)
}

func GetFloat(key string) (f float64, err error) {
	return GlobalConfig.GetFloat(key)
}