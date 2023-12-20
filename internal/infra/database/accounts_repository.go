package database

import "mongoapi/internal/domain/accounts"

type AccountsRepository struct{}

func (r *AccountsRepository) Create(a *accounts.Account) error {
	return nil
}

func (r *AccountsRepository) FindById(id string) (*accounts.Account, error) {
	return nil, nil
}
