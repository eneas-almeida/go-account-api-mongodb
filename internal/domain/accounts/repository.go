package accounts

type RepositoryInterface interface {
	Create(a *Account) error
	FindById(id string) (*Account, error)
}
