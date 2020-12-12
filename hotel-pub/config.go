package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (config *hotelOfferPubConfig) init() (err error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return errors.New("Error reading config: " + err.Error())
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return errors.New("Error processing config: " + err.Error())
	}
	return
}

func (config *hotelOfferPubConfig) setDefaults() {

	if config.RMqConfig.Address == "" {
		config.RMqConfig.Address = "127.0.0.1"
	}
	if config.RMqConfig.Port == "" {
		config.RMqConfig.Port = "5672"
	}
}
