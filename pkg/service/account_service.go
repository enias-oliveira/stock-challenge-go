package services

import (
	"context"
	"crypto/sha256"

	"stock-challenge-go/pkg/domain"

	repoInterface "stock-challenge-go/pkg/repository/interface"
	srvcInterface "stock-challenge-go/pkg/service/interface"

	"github.com/sethvargo/go-password/password"
)

type AccountService struct {
	accRepo repoInterface.AccountRepository
}

func NewAccountService(repo repoInterface.AccountRepository) srvcInterface.AccountService {
	return &AccountService{
		accRepo: repo,
	}
}

func (s *AccountService) Register(ctx context.Context, account domain.Account) (domain.Account, error) {
	newPassword, gpErr := password.Generate(32, 4, 4, false, false)

	account.Password = newPassword

	if gpErr != nil {
		return account, gpErr
	}

	hashedPassword := sha256.Sum256([]byte(newPassword))
	account.PasswordHash = hashedPassword[:]

	account, cuErr := s.accRepo.Save(ctx, account)

	if cuErr != nil {
		return account, cuErr
	}

	return account, nil
}
