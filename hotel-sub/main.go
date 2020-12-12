package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	configFile    = "./dependency/hotelOfferMgrConfig.yaml"
	isSqlInit     = true
	isRabbitInit  = true
	rmqConn       *amqp.Connection
	sqlConn       *gorm.DB
	config        hotelOfferMgrConfig
	TableConn     map[string]*gorm.DB
	startConsumer = true
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
	log.Debugf("Config: %v", config)
	sqlConn, err = sqlConnect(config.DbConfig)
	if err != nil {
		log.Errorf("Error [%s] connecting to database", err.Error())
		isSqlInit = false
	} else {
		log.Infof("DB Connection Initialized successfully.")
	}
	rmqConn, err = rmqConnect(config.RMqConfig)
	if err != nil {
		log.Errorf("Error [%s] connecting to rabbit MQ server.", err.Error())
		isRabbitInit = false
	} else {
		log.Infof("Rabbit MQ Connection Initialized successfully.")
	}
	TableConn = make(map[string]*gorm.DB)

}

func main() {
	var ch chan bool
	err := initDatabase(sqlConn)
	if err != nil {
		log.Errorf("Error initializing database")
		return
	}
	go startHotelStatsConsumer()
	if startConsumer {
		<-ch
	}
}

func generateDSN(config hotelOfferMgrConfig) (dsn string) {
	dbConfig := config.DbConfig
	return config.DbConfig.Username + ":" + config.DbConfig.Password + "@tcp(" +
		dbConfig.Address + ":" + dbConfig.Port + ")/" +
		dbConfig.DatabaseName
}

func sqlConnect(conf DbConfig) (conn *gorm.DB, err error) {
	dsn := generateDSN(config)
	log.Debugf("DNS: %s", dsn)
	conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
