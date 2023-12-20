package usecases

import "mongoapi/internal/domain/accounts"

type AccountsUseCase struct {
	Repository accounts.RepositoryInterface
}

func (u *AccountsUseCase) Create(d *accounts.InputAccountDto) (*accounts.OutputAccountDto, error) {
	println("AccountsUseCase.Create showwww!!!!")
	return nil, nil
}
