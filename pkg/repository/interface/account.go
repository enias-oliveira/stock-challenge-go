package interfaces

import (
	"context"
	"stock-challenge-go/pkg/domain"
)

type AccountRepository interface {
	Save(ctx context.Context, account domain.Account) (domain.Account, error)
	FindByEmail(ctx context.Context, email string) (domain.Account, error)
}
