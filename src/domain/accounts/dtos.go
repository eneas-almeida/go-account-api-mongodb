package accounts

import (
	"time"
)

type InputAccountDto struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type OutputAccountDto struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}
