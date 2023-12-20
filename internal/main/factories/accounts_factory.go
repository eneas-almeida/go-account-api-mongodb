package factories

import (
	"fmt"
	"mongoapi/internal/application/usecases"
	"mongoapi/internal/infra/database"
	"mongoapi/internal/presentation/handlers"
)

type AccountsFactory struct{}

func (f *AccountsFactory) MakeAccounts() {
	repository := database.AccountsRepository{}
	usecases := usecases.AccountsUseCase{Repository: &repository}
	handlers := handlers.AccountsHandler{UseCases: &usecases}

	fmt.Println(handlers)
}
