package interfaces

import (
	"stock-challenge-go/pkg/domain"
)

type AccountRepository interface {
	Save(domain.Account) (domain.Account, error)
	FindByEmail(string) (domain.Account, error)
}
