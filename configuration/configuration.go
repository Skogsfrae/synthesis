package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Host        string
	Port        string
	PostgresUri string `yaml:"postgres-uri"`
	KeyTimer    int    `yaml:"key-timer"`
}

var Config *Configuration

// todo: implement configuration from env
func LoadConfiguration(path string) {
	Config = &Configuration{
		Host:        "127.0.0.1",
		Port:        "3000",
		PostgresUri: "postgresql://localhost:5432/synthesis",
		KeyTimer:    600,
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Cannot open configuration file, loading default config")
		return
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		log.Fatalf("Cannot load configuration")
	}

	log.Println("Configuration loaded")
}
