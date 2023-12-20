package accounts

type RepositoryInterface interface {
	Create(account *Account) (string, error)
	FindById(id string) (*Account, error)
}
