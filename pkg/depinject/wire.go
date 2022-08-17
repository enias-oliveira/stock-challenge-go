//go:build wireinject
// +build wireinject

package depinject

import (
	"github.com/google/wire"

	client "stock-challenge-go/pkg/client"
	config "stock-challenge-go/pkg/config"
	db "stock-challenge-go/pkg/db"
	http "stock-challenge-go/pkg/http"
	handler "stock-challenge-go/pkg/http/handler"
	repository "stock-challenge-go/pkg/repository"
	service "stock-challenge-go/pkg/service"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewAccountRepository,
		repository.NewStockRepository,
		repository.NewHistoryRepository,
		client.NewStooqClient,
		service.NewAccountService,
		service.NewStockService,
		service.NewHistoryService,
		handler.NewAccountHandler,
		handler.NewStockHandler,
		handler.NewHistoryHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
