package factories

import (
	"mongoapi/internal/application/usecases/accounts"
	"mongoapi/internal/infra/database"
	"mongoapi/internal/presentation/handlers"
)

func MakeAccounts() handlers.AccountsHandler {
	repository := database.AccountsRepository{}
	usecases := accounts.AccountsUseCase{Repository: &repository}
	handlers := handlers.AccountsHandler{UseCases: &usecases}
	return handlers
}
