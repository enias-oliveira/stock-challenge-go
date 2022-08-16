package interfaces

import (
	"context"
	"stock-challenge-go/pkg/domain"
)

type AccountRepository interface {
	Save(ctx context.Context, account domain.Account) (domain.Account, error)
}
