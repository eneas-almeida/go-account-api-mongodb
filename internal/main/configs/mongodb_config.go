package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Uri            string
	DatabaseName   string
	CollectionName string
}

func NewMongoConfig(uri, database, collection string) (*MongoConfig, error) {
	return &MongoConfig{
		Uri:            uri,
		DatabaseName:   database,
		CollectionName: collection,
	}, nil
}

func (cfg *MongoConfig) client() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.Uri))

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (cfg *MongoConfig) Collection() *mongo.Collection {
	client, err := cfg.client()

	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	return client.Database(cfg.DatabaseName).Collection(cfg.CollectionName)
}

func (cfg *MongoConfig) Disconnect() {
	client, err := cfg.client()

	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	client.Disconnect(context.Background())
}
