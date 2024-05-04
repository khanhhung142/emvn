package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var config Config

func GetConfig() Config {
	return config
}

// Load the config from the config.yaml file
// In prod, the config.yaml file will be created by reading the environment variables and writing them to the file. More detail: entrypoint.sh
func InitConfig() {
	var c Config
	absPath, err := filepath.Abs("config/yaml/config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	yamlFile, err := os.ReadFile(absPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	config = c
}
