package config

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type Config struct {
	store map[string]string
}

var GlobalConfig Config

func init() {
	GlobalConfig.store = make(map[string]string)
}

// Load key-value pairs from ini-like config file
// key=value
// Any line, starting # - comment

func (this *Config)LoadFromFile(config_file string) (err error) {
	err = nil
	file, err := os.Open(config_file)
	if err != nil {
		log.Fatalln("Can not open config file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") && strings.Index(line, "=") != -1 {
			line_parts := strings.Split(line, "=")
			this.store[line_parts[0]] = line_parts[1]
		}
	}
	return
}
