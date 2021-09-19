package config

import (
	"log"
	"os"
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

func LoadEnv() (Env, error) {
	var env Env
	envName := os.Getenv("ENV")
	viper.AddConfigPath(".")
	if envName == "PROD" {
		log.Println("Load env from ", envName)
		viper.SetConfigName("prod")
	} else if envName == "DEV" {
		log.Println("Load env from ", envName)
		viper.SetConfigName("dev")
	} else {
		log.Println("Load envs directly from system env variables")
		duration, err := time.ParseDuration(os.Getenv("TOKEN_DURATION"))
		if err != nil {
			log.Panic(err)
		}
		env = Env{
			MongoDBUri:           os.Getenv("MONGODB_URI"),
			Host:                 os.Getenv("HOST"),
			DatabaseName:         os.Getenv("DATABASE_NAME"),
			SymmetricKey:         os.Getenv("SYMMETRIC_KEY"),
			TokenDuration:        duration,
			NetworkHost:          os.Getenv("NETWORK_HOST"),
			SmartContractAddress: os.Getenv("SMART_CONTRACT_ADDRESS"),
		}
		log.Println(env)
		return env, nil
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = viper.Unmarshal(&env)
	return env, err
}
