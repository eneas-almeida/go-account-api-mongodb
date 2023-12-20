package accounts

type RepositoryInterface interface {
	Create(account *Account) (string, error)
	FindOneById(id string) (*Account, error)
}
