package config

import (
	"errors"
	"os"
	"strings"
)

//Appconfig contains the application's configuration
type Appconfig struct {
	KafkaBrokers []string
	KafKaTopic   string
}

//GetAppConfiguration returns the configuration from environmental variables
func GetAppConfiguration() (Appconfig, error) {

	var appConfig = Appconfig{}

	var topic = os.Getenv("TOPIC")

	if len(topic) == 0 {
		return appConfig, errors.New("kafka topic hasn't been specified.")
	}

	var brokers = os.Getenv("BROKERS")

	if len(brokers) == 0 {
		return appConfig, errors.New("kafka brokers haven't been specified.")
	}

	appConfig.KafKaTopic = topic
	appConfig.KafkaBrokers = strings.Split(brokers, ",")

	return appConfig, nil

}
