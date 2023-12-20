package accounts

import "mongoapi/internal/domain/accounts"

type AccountsUseCase struct {
	Repository accounts.RepositoryInterface
}

func (u *AccountsUseCase) Create(d *accounts.InputAccountDto) (*accounts.OutputAccountDto, error) {
	account, err := accounts.NewAccount(d.ID, d.Name, d.Email, d.Password)

	if err != nil {
		return nil, err
	}

	// err = u.Repository.Create(account)

	// if err != nil {
	// 	return nil, err
	// }

	output := &accounts.OutputAccountDto{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		CreatedAt: account.CreatedAt.UTC(),
	}

	return output, nil
}
