package database

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var (
	once     sync.Once
	Instance *Database
)

func SetupMongoDB(DBURI, DBName string) {
	once.Do(func() {
		mongoOptions := options.Client().ApplyURI(DBURI)
		client, err := mongo.Connect(context.TODO(), mongoOptions)
		if err != nil {
			log.Panic("Error connecting to MongoDB: ", err)
		}

		db := client.Database(DBName)

		Instance = &Database{
			Client: client,
			DB:     db,
		}

		defer func() {
			if err = client.Disconnect(context.Background()); err != nil {
				log.Panic("Error disconnecting from MongoDB: ", err)
			}
		}()
	})
}
