package factories

import (
	"log"
	"mongoapi/internal/main/configs"
)

func MakeMongoConnection() *configs.MongoConfig {
	uri := "mongodb://customers:customers@localhost:27037/customers"
	database := "customers"
	collection := "accounts"

	connection, err := configs.NewMongoConfig(uri, database, collection)

	if err != nil {
		log.Fatal("Error creating repository:", err)
	}

	return connection
}
