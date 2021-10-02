package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	API_HOST          string
	PAWN_PATH         string
	BID_PATH          string
	BID_PAWN_PATH     string
	NOTIFY_HOST       string
	NOTIFICATION_PATH string
	CONTRACT_ADDRESS  string
	NETWORK_ADDRESS   string
	RABBIT_MQ_URI     string
}

func LoadEnv() (*Env, error) {
	var env Env
	viper.AddConfigPath(".")
	log.Println("Environment: ", os.Getenv("ENV"))
	if os.Getenv("ENV") == "PROD" {
		viper.SetConfigName("prod")
	} else if os.Getenv("ENV") == "DEV" {
		viper.SetConfigName("dev")
		log.Println("load env config from dev.env")
	} else {
		log.Print("load variable from environment")
		env = Env{
			API_HOST:          os.Getenv("API_HOST"),
			NOTIFY_HOST:       os.Getenv("NOTIFY_HOST"),
			CONTRACT_ADDRESS:  os.Getenv("CONTRACT_ADDRESS"),
			NETWORK_ADDRESS:   os.Getenv("NETWORK_ADDRESS"),
			PAWN_PATH:         os.Getenv("PAWN_PATH"),
			BID_PATH:          os.Getenv("BID_PATH"),
			BID_PAWN_PATH:     os.Getenv("BID_PAWN_PATH"),
			NOTIFICATION_PATH: os.Getenv("NOTIFICATION_PATH"),
			RABBIT_MQ_URI:     os.Getenv("RABBIT_MQ_URI"),
		}
		return &env, nil
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = viper.Unmarshal(&env)
	return &env, err
}
