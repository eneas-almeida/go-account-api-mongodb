package accounts

import "mongoapi/internal/domain/accounts"

type AccountsUseCase struct {
	Repository accounts.RepositoryInterface
}
