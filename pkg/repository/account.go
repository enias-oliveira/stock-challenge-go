package repository

import (
	"context"

	"gorm.io/gorm"

	"stock-challenge-go/pkg/domain"
	repoInterface "stock-challenge-go/pkg/repository/interface"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repoInterface.AccountRepository {
	return &accountRepository{db}
}

func (ar *accountRepository) Save(ctx context.Context, account domain.Account) (domain.Account, error) {
	err := ar.db.Create(&account).Error

	return account, err
}
