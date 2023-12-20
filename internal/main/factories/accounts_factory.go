package factories

import (
	"mongoapi/internal/application/usecases/accounts"
	"mongoapi/internal/infra/database/repositories"
	"mongoapi/internal/presentation/handlers"
)

func MakeAccountsHandler() handlers.AccountsHandler {
	repository := repositories.AccountsRepository{Connection: MakeMongoConnection()}

	usecases := accounts.AccountsUseCase{Repository: &repository}

	handlers := handlers.AccountsHandler{UseCases: &usecases}

	return handlers
}
