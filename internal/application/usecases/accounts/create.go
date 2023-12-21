package accounts

import "mongoapi/internal/domain/accounts"

func (u *AccountsUseCase) Create(d *accounts.InputAccountDto) (*accounts.OutputAccountDto, error) {
	account, err := accounts.NewAccount(d.Name, d.Email, d.Password)

	if err != nil {
		return nil, err
	}

	id, err := u.Repository.Create(account)

	if err != nil {
		return nil, err
	}

	output := &accounts.OutputAccountDto{
		ID:        id,
		Name:      account.Name,
		Email:     account.Email,
		CreatedAt: account.CreatedAt.UTC(),
	}

	return output, nil
}
