package main

import (
	"log"

	"github.com/uss-kelvin/NFTPawningShopBackend/server"
	"github.com/uss-kelvin/NFTPawningShopBackend/server/config"
)

func main() {
	connection, err := config.GetConnection()
	if err != nil {
		log.Panic(err)
	}
	err = connection.TestConnection()
	if err != nil {
		log.Panic(err)
	}
	env, err := config.LoadEnv("")
	if err != nil {
		log.Panic(err)
	}
	app, err := server.NewServer(connection, &env)
	if err != nil {
		log.Panic(err)
	}
	if err = app.Start(env.Host); err != nil {
		log.Panic(err)
	}
	log.Printf("Server is running at %v \n", env.Host)
}
