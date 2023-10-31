package data

import (
	"UserAPI/internal/api/utils"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func NewDatabaseConfig() utils.DatabaseConfig {
	var config utils.DatabaseConfig

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
