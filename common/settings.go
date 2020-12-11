package common

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Server, MongoDBHost, DBUser, DBPwd, DataBase string
}

// AppConfig object would be used for storing Configuration settings
// from config.json
var AppConfig configuration

func initConfig() {
	loadAppConfig()

}

func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("Error loading config.json %s\n", err)
	}

	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[Decoding AppConfig] %s\n", err)
	}
	fmt.Println("Appconfig: ", AppConfig)
}
