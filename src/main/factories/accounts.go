package factories

import (
	"mongoapi/src/application/usecases/accounts"
	"mongoapi/src/infra/database/repositories"
	"mongoapi/src/presentation/handlers"
)

func MakeAccountsHandler() handlers.AccountsHandler {
	repository := repositories.AccountsRepository{}
	usecases := accounts.AccountsUseCase{Repository: &repository}
	handlers := handlers.AccountsHandler{UseCases: &usecases}
	return handlers
}
