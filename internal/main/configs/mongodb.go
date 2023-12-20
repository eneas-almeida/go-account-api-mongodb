package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetMongoDbConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27037"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func GetMongoDbCollection(dbName string, collectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return collection, nil
}
