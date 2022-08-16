package interfaces

import (
	"context"

	domain "stock-challenge-go/pkg/domain"
)

type AccountService interface {
	Register(ctx context.Context, account domain.Account) (domain.Account, error)
}
