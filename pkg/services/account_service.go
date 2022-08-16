package services

import (
	"context"
	"crypto/sha256"

	"stock-challenge-go/pkg/domain"
	accService "stock-challenge-go/pkg/services/interface"

	"github.com/sethvargo/go-password/password"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type accountService struct{}

func NewAccountService() accService.AccountService {
	return &accountService{}
}

func (s *accountService) Register(ctx context.Context, account domain.Account) (domain.Account, error) {
	dsn := "root@/stock-challenge"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&domain.Account{})

	if err != nil {
		panic(err)
	}

	newPassword, gpErr := password.Generate(32, 4, 4, false, false)

	account.Password = newPassword

	if gpErr != nil {
		return account, gpErr
	}

	hashedPassword := sha256.Sum256([]byte(newPassword))
	account.PasswordHash = hashedPassword[:]

	cuErr := db.Create(&account).Error

	if cuErr != nil {
		return account, cuErr
	}

	return account, nil
}
