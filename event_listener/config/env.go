package config

import (
	"log"

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
}

func LoadEnv() (*Env, error) {
	var env Env
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = viper.Unmarshal(&env)
	return &env, err
}
