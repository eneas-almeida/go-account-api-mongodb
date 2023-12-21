package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConfig struct {
	Client         *mongo.Client
	Uri            string
	DatabaseName   string
	CollectionName string
}

func NewMongoConfig(uri, database, collection string) (*MongoConfig, error) {
	options := options.Client().ApplyURI(uri).SetMaxPoolSize(6)

	client, err := mongo.Connect(context.Background(), options)

	if err != nil {
		return nil, err
	}

	return &MongoConfig{
		Uri:            uri,
		Client:         client,
		DatabaseName:   database,
		CollectionName: collection,
	}, nil
}

func (c *MongoConfig) setConnection() {
	err := c.Client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.Uri))

		if err != nil {
			log.Fatal(err)
		}

		c.Client = client
	}
}

func (c *MongoConfig) Collection() *mongo.Collection {
	c.setConnection()

	return c.Client.Database(c.DatabaseName).Collection(c.CollectionName)
}

func (c *MongoConfig) Disconnect() {
	c.Client.Disconnect(context.Background())
}
