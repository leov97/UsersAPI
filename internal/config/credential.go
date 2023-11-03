package config

import (
	"UserAPI/internal/models"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func NewDatabaseConfig() models.DatabaseConfig {
	var config models.DatabaseConfig

	dataconfig, err := os.Open("config.yaml")
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
