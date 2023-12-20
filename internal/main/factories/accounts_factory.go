package factories

import (
	"log"
	"mongoapi/internal/application/usecases/accounts"
	"mongoapi/internal/infra/database"
	"mongoapi/internal/presentation/handlers"
)

func MakeAccounts() handlers.AccountsHandler {
	client := "mongodb://customers:customers@localhost:27037/customers"
	databaseName := "customers"
	collectionName := "accounts"

	repository, err := database.NewAccountsRepository(client, databaseName, collectionName)

	if err != nil {
		log.Fatal("Error creating repository:", err)
	}

	defer repository.Disconnect()

	usecases := accounts.AccountsUseCase{Repository: repository}
	handlers := handlers.AccountsHandler{UseCases: &usecases}

	return handlers
}
