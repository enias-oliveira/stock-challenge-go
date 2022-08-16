//go:build wireinject
// +build wireinject

package depinject

import (
	"github.com/google/wire"

	config "stock-challenge-go/pkg/config"
	db "stock-challenge-go/pkg/db"
	http "stock-challenge-go/pkg/http"
	repository "stock-challenge-go/pkg/repository"
	service "stock-challenge-go/pkg/service"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewAccountRepository, service.NewAccountService, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
