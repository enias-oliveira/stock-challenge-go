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

func (ar *accountRepository) Save(account domain.Account) (domain.Account, error) {
	err := ar.db.Create(&account).Error

	return account, err
}

func (ar *accountRepository) FindByEmail(ctx context.Context, email string) (domain.Account, error) {
	account := domain.Account{
		Email: email,
	}
	err := ar.db.Where(&account, "email").Find(&account).Error

	return account, err
}
