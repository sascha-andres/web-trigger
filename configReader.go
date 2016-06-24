package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func readConfig(fileToRead string) (configDocument, error) {
	var config configDocument

	// check if file exists
	if _, err := os.Stat(fileToRead); os.IsNotExist(err) {
		return config, err
	}

	data, err := ioutil.ReadFile(fileToRead)
	if err != nil {
		return config, err
	}

	// at this point we have the data let's try to deserialize
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config, nil
}
