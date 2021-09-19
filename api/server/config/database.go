package config

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	client *mongo.Client
}

func GetConnection() (*Connection, error) {
	env, err := LoadEnv()
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

func (con *Connection) GetSession(callback func(*gin.Context, mongo.SessionContext) error) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var session mongo.Session
		var err error
		if session, err = con.client.StartSession(); err != nil {
			log.Panic(err)
		}
		if err = session.StartTransaction(); err != nil {
			log.Panic(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
			// execute call back here
			log.Println("execute callback")
			err := callback(c, sc)
			if err != nil {
				log.Println("error happens right here")
				session.AbortTransaction(sc)
				return err
			}
			log.Println("go here :)")
			if err = session.CommitTransaction(sc); err != nil {
				log.Panic(err)
				return err
			}
			return nil
		}); err != nil {
			log.Panic(err)
		}
		session.EndSession(ctx)
	}
	return fn
}
