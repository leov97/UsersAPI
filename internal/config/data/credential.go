package data

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Network  string `yaml:"network"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"dbName"`
	} `yaml:"database"`
}

func NewDatabaseConfig() DatabaseConfig {
	var config DatabaseConfig

	dataconfig, err := os.Open("cred.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer dataconfig.Close()

	contents, err := ioutil.ReadAll(dataconfig)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
