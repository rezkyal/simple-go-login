package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rezkyal/simple-go-login/entity/config"
)

func InitConfig() (*config.Config, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("files/config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, fmt.Errorf("[InitConfig] error when opening config file, err: %+v", err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result *config.Config = &config.Config{}
	err = json.Unmarshal([]byte(byteValue), result)

	if err != nil {
		return nil, fmt.Errorf("[InitConfig] error when unmarshall, err: %+v", err)
	}

	return result, nil
}
