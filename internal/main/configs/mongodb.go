package configs

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(collection string) *mongo.Collection {
	return db.Collection(collection)
}

func InitDB() error {
	// uri := os.Getenv("MONGODB_URI")

	// if uri == "" {
	// 	return errors.New("you must set your 'MONGODB_URI' environmental variable")
	// }

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://customers:customers@localhost:27037/customers"))

	if err != nil {
		return err
	}

	db = client.Database("customers")

	return nil
}

func CloseDB() error {
	return db.Client().Disconnect(context.Background())
}
