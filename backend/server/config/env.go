package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Env struct {
	MongoDBUri           string        `mapstructure:"MongoDB_URI"`
	Host                 string        `mapstructure:"HOST"`
	DatabaseName         string        `mapstructure:"DATABASE_NAME"`
	SymmetricKey         string        `mapstructure:"SYMMETRIC_KEY"`
	TokenDuration        time.Duration `mapstructure:"TOKEN_DURATION"`
	NetworkHost          string        `mapstructure:"NETWORK_HOST"`
	SmartContractAddress string        `mapstructure:"SMART_CONTRACT_ADDRESS"`
}

func LoadEnv(path string) (Env, error) {
	var env Env
	if path == "" {
		path = "."
	}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = viper.Unmarshal(&env)
	return env, err
}
