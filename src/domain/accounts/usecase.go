package accounts

type UseCaseInterface interface {
	Create(d *InputAccountDto) (*OutputAccountDto, error)
	FindOneById(id string) (*OutputAccountDto, error)
}
