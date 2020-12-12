package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	configFile   = "./dependency/hotelOfferPubConfig.yaml"
	config       hotelOfferPubConfig
	rmqConn      *amqp.Connection
	isRabbitInit = true
)

func init() {
	var err error
	log.SetLevel(log.DebugLevel)
	err = config.init()
	if err != nil {
		log.Errorf("Error initializing from config file %s", configFile)
		log.Infof("Will attempt with defaults")
	}
	config.setDefaults()
	rmqConn, err = rmqConnect(config.RMqConfig)
	if err != nil {
		log.Errorf("Error [%s] connecting to rabbit MQ server.", err.Error())
		isRabbitInit = false
	} else {
		log.Infof("Rabbit MQ Connection Initialized successfully.")
	}
}

func main() {
	ch, err := rmqConn.Channel()
	q, err := ch.QueueDeclare("hotelStats", false, false, false, false, nil)
	if err != nil {
		log.Errorf("Error [%s] when declaring queue", err.Error())
		return
	}
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(hotelData),
	})
}

func rmqConnect(conf RMqConfig) (conn *amqp.Connection, err error) {
	url := "amqp://" + conf.Address + ":" + conf.Port
	log.Debugf("Connecting to Rabbit MQ [%s]", url)
	conn, err = amqp.Dial(url)
	return
}
