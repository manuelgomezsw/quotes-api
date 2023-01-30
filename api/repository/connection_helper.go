package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	ConnectionString = "mongodb+srv://dbuser:Pergamino18.@quotes-cluster.0lj6f.mongodb.net/test"
	DB               = "quotes-db"
	COLLECTION       = "quotes"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(ConnectionString)

		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceError
}
