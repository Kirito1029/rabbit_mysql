package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (config *hotelOfferMgrConfig) init() (err error) {
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

func (config *hotelOfferMgrConfig) setDefaults() {
	if config.DbConfig.Address == "" {
		config.DbConfig.Address = "127.0.0.1"
	}
	if config.DbConfig.Port == "" {
		config.DbConfig.Port = "3306"
	}
	if config.DbConfig.DatabaseName == "" {
		config.DbConfig.DatabaseName = "hoteloffer"
	}
	if config.RMqConfig.Address == "" {
		config.RMqConfig.Address = "127.0.0.1"
	}
	if config.RMqConfig.Port == "" {
		config.RMqConfig.Port = "5672"
	}
	if config.DbConfig.Username == "" {
		config.DbConfig.Username = "root"
	}
	if config.DbConfig.Password == "" {
		config.DbConfig.Password = "password"
	}
}
