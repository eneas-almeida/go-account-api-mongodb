package accounts

import (
	"errors"
	"mongoapi/internal/domain/shared"
	"time"
)

type Account struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewAccount(id string, name string, email string, password string) (*Account, error) {
	account := &Account{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	account.prepare()

	return account, nil
}

func (a *Account) prepare() {
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Account) isValid() error {
	notification := shared.CreateNotification("account")

	if !Id(a.ID) {
		notification.AddError("Id is invalid")
	}

	if !Name(a.Name) {
		notification.AddError("Name is invalid")
	}

	if !Email(a.Email) {
		notification.AddError("Email is invalid")
	}

	if !Password(a.Password) {
		notification.AddError("Password is invalid")
	}

	if notification.HasError() {
		return errors.New(notification.GetErrors())
	}

	return nil
}
