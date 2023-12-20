package database

import (
	"context"
	"log"
	"mongoapi/internal/domain/accounts"
	"mongoapi/internal/main/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountsRepository struct {
	Client         *mongo.Client
	DatabaseName   string
	CollectionName string
}

func NewAccountsRepository(uri, databaseName, collectionName string) (*AccountsRepository, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	return &AccountsRepository{
		Client:         client,
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}, nil
}

func (repo *AccountsRepository) Disconnect() {
	if repo.Client != nil {
		err := repo.Client.Disconnect(context.Background())

		if err != nil {
			log.Println("Error disconnecting from MongoDB:", err)
		}
	}
}

func (repo *AccountsRepository) Create(account *accounts.Account) (string, error) {
	collection, err := configs.GetMongoDbConnection("customers", "accounts")

	if err != nil {
		return "", err
	}

	doc := bson.M{
		"name":       account.Name,
		"email":      account.Email,
		"password":   account.Password,
		"created_at": account.CreatedAt,
		"updated_at": account.UpdatedAt,
	}

	if err != nil {
		return "", err
	}

	res, err := collection.InsertOne(context.Background(), doc)

	if err != nil {
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (repo *AccountsRepository) FindById(id string) (*accounts.Account, error) {
	return nil, nil
}
