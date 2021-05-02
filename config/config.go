package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadConfiguration() Configuration {

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	jsonContent, err := ioutil.ReadFile(fmt.Sprintf("%s/spotmicro.json", cwd))
	if err != nil {
		log.Fatal(err)
	}
	var configuration Configuration
	err = json.Unmarshal(jsonContent, &configuration)
	if err != nil {
		fmt.Println("Cannot read config file")
		os.Exit(1)
	}
	return configuration
}
