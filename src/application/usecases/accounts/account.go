package accounts

import "mongoapi/src/domain/accounts"

type AccountsUseCase struct {
	Repository accounts.RepositoryInterface
}

func NewAccountsUseCase(repository accounts.RepositoryInterface) *AccountsUseCase {
	return &AccountsUseCase{
		Repository: repository,
	}
}
