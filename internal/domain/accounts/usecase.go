package accounts

type UseCaseInterface interface {
	Create(d *InputAccountDto) (*OutputAccountDto, error)
}
