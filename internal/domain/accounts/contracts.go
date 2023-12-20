package accounts

type InputAccountDto struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type OutputAccountDto struct {
	ID    string
	Name  string
	Email string
}

type RepositoryInterface interface {
	Create(a *Account) error
	FindById(id string) (*Account, error)
}

type UseCasesInterface interface {
	Create(d *InputAccountDto) (*OutputAccountDto, error)
}
