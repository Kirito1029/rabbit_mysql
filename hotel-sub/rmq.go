package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func rmqConnect(conf RMqConfig) (conn *amqp.Connection, err error) {
	url := "amqp://" + conf.Address + ":" + conf.Port
	log.Debugf("Connecting to Rabbit MQ [%s]", url)
	conn, err = amqp.Dial(url)
	return
}

func startHotelStatsConsumer() {
	ch, err := rmqConn.Channel()
	if err != nil {
		log.Errorf("Error [%s] creating channel", err.Error())
		return
	}
	defer ch.Close()
	hotelQueue, err := ch.QueueDeclare("hotelStats", false, false, false, false, nil)
	if err != nil {
		log.Errorf("Error [%s] when declaring queue", err.Error())
		return
	}

	if !startConsumer {
		return
	}
	msgs, err := ch.Consume(hotelQueue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Errorf("Error [%s] registering consumer", err.Error())
		return
	}
	for d := range msgs {
		var offerList HotelOffers
		err := json.Unmarshal(d.Body, &offerList)
		if err != nil {
			log.Errorf("Error [%s] unmarshalling message [%s]", err.Error(), "d.Body")
			d.Ack(false)
			continue
		}
		offerList.processHotelOffer()
		d.Ack(false)
	}
}
