//go:build wireinject
// +build wireinject

package depinject

import (
	"github.com/google/wire"

	config "stock-challenge-go/pkg/config"
	db "stock-challenge-go/pkg/db"
	http "stock-challenge-go/pkg/http"
	handler "stock-challenge-go/pkg/http/handler"
	repository "stock-challenge-go/pkg/repository"
	service "stock-challenge-go/pkg/service"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewAccountRepository, service.NewAccountService, handler.NewAccountHandler, handler.NewStockHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
