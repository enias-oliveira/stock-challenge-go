package interfaces

import (
	domain "stock-challenge-go/pkg/domain"
)

type AccountService interface {
	Register(domain.Account) (domain.Account, error)
	ValidateAccount(domain.Account) (domain.Account, error)
}
