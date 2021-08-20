package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	client *mongo.Client
}

func GetConnection() (*Connection, error) {
	env, err := LoadEnv("")
	if err != nil {
		log.Panic(err)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(env.MongoDBUri))
	if err != nil {
		log.Panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	connection := &Connection{
		client: client,
	}
	return connection, err
}

func (d *Connection) GetDatabase(databaseName string) *mongo.Database {
	return d.client.Database(databaseName)
}

func (d *Connection) TestConnection() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := d.client.Ping(ctx, readpref.Primary())
	return err
}

func (c *Connection) GetDatabaseNames() ([]string, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return c.client.ListDatabaseNames(ctx, bson.M{})
}
